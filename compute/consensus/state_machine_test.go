package consensus

import (
	"compute/model"
	"encoding/json"
	"fmt"
	"sync"
	"testing"
	"time"
)

// state_machine_test.go

func TestMarshal(t *testing.T) {
	data := newFsm()
	data.Pointer = 111
	data.Queue = append(data.Queue, value{
		Site: 1,
		Value: model.ConsensusValue{
			Type:  model.Success,
			Value: 1,
		},
	})
	//ret, err := json.Marshal(data.Queue)
	ret, err := json.Marshal(data)
	t.Log(string(ret))
	t.Log(err)
	var d fsm
	err = json.Unmarshal(ret, &d)
	t.Log(d)
}

type queue struct {
	data    []int
	idx     int
	watcher chan int
	sync.Mutex
}

func (q *queue) push(data int) {
	fmt.Println("push", data)
	q.Lock()
	q.data = append(q.data, data)
	if q.watcher != nil {
		q.watcher <- data
		q.idx = len(q.data)
		fmt.Println("当前idx", q.idx)
		//q.idx++
		fmt.Println("idx=", q.idx)
		q.watcher = nil
	}
	q.Unlock()
}

func (q *queue) watch(target int, waiter chan int) {
	for i := q.idx; i < len(q.data); i++ {
		fmt.Println("当前位置", i)
		if q.data[i] == target {
			q.idx++
			waiter <- q.data[i]
			return
		}
	}
	fmt.Println("没有watch到值,等待push")
	q.watcher = waiter
}

func TestPointer(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	var q queue
	go func() {
		defer wg.Done()
		w := make(chan int, 1)
		// goroutine 1
		q.push(1)
		q.push(11)
		q.push(111)
		q.watch(2, w)
		fmt.Println("goroutine 1 watch 2", <-w)

	}()
	go func() {
		defer wg.Done()
		// goroutine 2
		time.Sleep(time.Second * 2)
		w := make(chan int, 1)
		q.watch(1, w)
		fmt.Println("goroutine 2 watch 1", <-w)
		w1 := make(chan int, 1)
		q.watch(11, w1)
		fmt.Println("goroutine 2 watch 11", <-w1)
		w2 := make(chan int, 1)
		q.watch(111, w2)
		fmt.Println("goroutine 2 watch 111", <-w2)
		q.push(2)
	}()
	wg.Wait()
}
