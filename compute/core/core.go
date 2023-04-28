package core

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"sync"
	"time"

	"compute/api/proto/container"
	"compute/api/proto/node"
	"compute/client"
	cfg "compute/config"
	"compute/consensus"
	"compute/db"
	zlog "compute/log"
	"compute/model"

	"github.com/imroc/req/v3"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var C = NewCore()

type Core struct {
	processTable *processTable
	lock         sync.Mutex
	node.UnimplementedNodeServiceServer
	container.UnimplementedContainerServiceServer
}

func NewCore() *Core {
	return &Core{
		processTable: newPT(),
	}
}

// StartProcess 启动进程,准备pid,建立共识节点,建立共识,启动容器发起计算
func (c *Core) StartProcess(pid model.Pid, funcId string, members []string, callback chan<- *model.Output, startErrCh chan<- error) {
	c.lock.Lock()
	defer c.lock.Unlock()
	p := new(pcb)
	p.callback = callback
	if len(members) == 1 {
		// 只有一个成员,单节点计算,不需要建立共识,同时容器内直接执行即可
		goto label
	} else {
		// 多节点参与计算
		r, err := consensus.NewRaft(pid, members)
		if err != nil {
			startErrCh <- err
		}
		p.consensus = r
		if err = p.consensus.BuildConsensus(); err != nil {
			startErrCh <- err
		}
	}
label:
	if m, err := client.GenerateTempFile(funcId); err != nil {
		startErrCh <- err
	} else {
		p.tempFilePath = m[0]
		p.fnName = m[1]
	}
	startErrCh <- nil
	// 将当前进程插入进程表
	c.processTable.put(pid, p)
	go func() {
		errorCh := make(chan error, 1)
		c.startContainer(pid, members, errorCh)
		timeOut := time.After(time.Second * 30)
		select {
		case <-timeOut:
			zlog.Info(TimeoutErr)
			c.processTable.delete(pid)
			_ = p.shutdown()
		case err := <-errorCh:
			if err != nil {
				zlog.Error("start podman failed", zap.Error(err))
				c.processTable.delete(pid)
				_ = p.shutdown()
			}
		}
	}()
	return
}

var isRm = flag.Bool("rm", true, "执行完之后是否删除容器")

const (
	// Go podman镜像中将go链接到了/bin/中
	Go        = "Go"
	imageName = "kamier:1.0"
)

// startContainer 启动容器执行计算方法
func (c *Core) startContainer(pid model.Pid, members []string, errCh chan error) {
	flag.Parse()
	zlog.Debug("start container...")
	p, ok := c.processTable.get(pid)
	if !ok {
		zlog.Error(PidNotExistsErr)
		errCh <- errors.New("not")
		return
	}
	pwd, err := os.Getwd()
	if err != nil {
		errCh <- err
		return
	}
	var cmdArgs []string
	if *isRm {
		cmdArgs = []string{"run", "--rm", "--network=host"}
	} else {
		cmdArgs = []string{"run", "--network=host"}
	}
	cmdArgs = append(cmdArgs,
		"-e", "SocketPath="+cfg.Cfg.SocketPath,
		"-e", "SelfName="+cfg.Cfg.NodeName,
		"-e", "NodeName="+pid.NodeName,
		"-e", "Serial="+strconv.FormatInt(pid.Serial, 10),
		"-e", "Host_IP="+cfg.Cfg.LocalAddr,
		"-e", "MembersLength="+strconv.Itoa(len(members)),
		"-e", "FnName="+p.fnName,
		// 挂载的文件
		"-v", "../containers/:/root/containers/",
		// sockets文件
		"-v", pwd+"/socket.sock:/root/containers/socket.sock",
		"-v", "./logs:/root/containers/logs",
		"-w", "/root/containers/",
		imageName,
		Go, "run", ".",
		//Go, "run", "main.go",
	)
	cmd := exec.Command("podman", cmdArgs...)
	//cmd = exec.Command("podman", "run", "-v ../container:/root/container", "-w /root/container", "kamier:1.0", "Go run main.go ")
	fmt.Println("[podman arg...]", cmd.String())
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err = cmd.Start(); err != nil {
		zlog.Error("container start failed", zap.Error(err))
		errCh <- err
		return
	}
	zlog.Debug("waiting...")
	if err = cmd.Wait(); err != nil {
		zlog.Error("exec failed", zap.String("stderr", stderr.String()))
		errCh <- errors.New(stderr.String())
		return
	}
	zlog.Debug("start exec")
	if err = c.finish(pid, model.Output{
		StdOut: stdout.String(),
		StdErr: stderr.String(),
	}); err != nil {
		zlog.Error("exec failed")
		errCh <- errors.New(stderr.String())
		return
	}
	errCh <- nil
}

// finish 执行完之后的结果处理
func (c *Core) finish(pid model.Pid, output model.Output) error {
	var err error
	defer c.processTable.delete(pid)
	defer func() {
		_, err = req.R().SetBodyJsonMarshal(pid).Delete(cfg.Cfg.StorageUrl + "/consensus/")
		if err != nil {
			zlog.Error("delete port num failed", zap.Error(err))
			return
		}
	}()
	zlog.Info("container exec success", zap.Any("output", output))
	if err = db.StoreOutput(pid, output); err != nil {
		return err
	}
	p, ok := c.processTable.get(pid)
	// 当前进程控制块存在并且callback不为nil(异步调用时callback为nil)
	if ok {
		if p.callback != nil {
			p.callback <- &output
			if err = p.shutdown(); err != nil {
				zlog.Warn("pcb shutdown failed", zap.Error(err))
				return err
			}
			return nil
		}
	}
	return errors.New("process not found")
}
