package core

import (
	"bytes"
	"compute/api/proto/container"
	"flag"
	"os"
	"os/exec"
	"strconv"
	"sync"

	"compute/api/proto/node"
	"compute/client"
	cfg "compute/config"
	"compute/consensus"
	"compute/db"
	zlog "compute/log"
	"compute/model"

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
func (c *Core) StartProcess(pid model.Pid, funcId string, members []string, callback chan<- *model.Output, errCh chan error) {
	c.lock.Lock()
	defer c.lock.Unlock()
	p := new(pcb)
	p.callback = callback
	if len(members) == 1 {
		// 只有一个成员,单节点计算,不需要建立共识,同时容器内直接执行即可
		goto label
	} else {
		// 多节点参与计算
		r, err := consensus.NewRaft(members)
		if err != nil {
			errCh <- err
		}
		p.consensus = r
		if err = r.BuildConsensus(); err != nil {
			errCh <- err
		}
	}
label:
	if err := client.GenerateTempFile(funcId); err != nil {
		errCh <- err
	}
	// 将当前进程插入进程表
	c.processTable.put(pid, p)
	go c.startContainer(pid, errCh)
}

var isRm = flag.Bool("rm", true, "执行完之后是否删除容器")

const (
	// Go podman镜像中将go链接到了/bin/中
	Go            = "Go"
	containerName = "kamier"
)

// startContainer 启动容器执行计算方法
func (c *Core) startContainer(pid model.Pid, errCh chan error) {
	flag.Parse()
	zlog.Debug("start container...")
	p, ok := c.processTable.get(pid)
	if !ok {
		errCh <- errors.New("not")
	}
	pwd, err := os.Getwd()
	if err != nil {
		errCh <- err
		return
	}
	var cmdArgs []string
	if *isRm {
		cmdArgs = []string{"run", "--network=host"}
	} else {
		cmdArgs = []string{"run", "--rm", "--network=host"}
	}
	cmdArgs = append(cmdArgs,
		"-e", "SocketPath", cfg.Cfg.SocketPath,
		"-e", "NodeName", cfg.Cfg.NodeName,
		"-e", "Serial", strconv.FormatInt(pid.Serial, 10),
		"-e", "Host", cfg.Cfg.GrpcAddr,
		// 挂载的文件
		"-v", "",
		// sockets文件
		"-v", pwd+"socket.sock:/",
		// TODO 工作目录
		"-w", "",
		containerName,
		Go, "run", "main.go",
	)
	cmd := exec.Command("podman", cmdArgs...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err = cmd.Start(); err != nil {
		zlog.Error("container start failed", zap.Error(err))
		errCh <- err
		return
	}
	if err = cmd.Wait(); err != nil {
		zlog.Error("exec failed", zap.Error(err))
		errCh <- err
		return
	}
	zlog.Debug("exec success")
	if err = c.finish(pid, model.Output{
		StdOut: stdout.String(),
		StdErr: stderr.String(),
	}); err != nil {
		errCh <- errors.New(stderr.String())
		return
	}
	errCh <- nil
}

// finish 执行完之后的结果处理
func (c *Core) finish(pid model.Pid, output model.Output) error {
	defer c.processTable.delete(pid)
	zlog.Info("container exec success")
	if err := db.StoreOutput(pid, output); err != nil {
		return err
	}
	p, ok := c.processTable.get(pid)
	// 当前进程控制块存在并且callback不为nil(异步调用时callback为nil)
	if ok && p.callback != nil {
		p.callback <- &output
		if err := p.shutdown(); err != nil {
			zlog.Warn("pcb shutdown failed", zap.Error(err))
			return err
		}
		return nil
	}
	return errors.New("process not found")
}
