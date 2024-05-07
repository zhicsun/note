package paradigm

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"testing"
	"time"
)

func TestWorkPoll(t *testing.T) {
	wm := NewWorkerManager(10)
	wm.StartWorkerPool()
}

type WorkerManager struct {
	workerChan chan *worker
	nWorkers   int
	sig        chan os.Signal
	exit       bool
}

func NewWorkerManager(workers int) *WorkerManager {
	return &WorkerManager{
		nWorkers:   workers,
		workerChan: make(chan *worker, workers),
		sig:        make(chan os.Signal, 1),
	}
}

func (wm *WorkerManager) StartWorkerPool() {
	signal.Notify(wm.sig, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-wm.sig
		wm.exit = true
		close(wm.workerChan)
	}()

	for i := 0; i < wm.nWorkers; i++ {
		wk := &worker{id: i}
		go wk.work(wm.workerChan)
	}

	wm.KeepLiveWorkers()
}

func (wm *WorkerManager) KeepLiveWorkers() {
	for wk := range wm.workerChan {
		fmt.Printf("Worker %d stopped with err: [%v] \n", wk.id, wk.err)
		wk.err = nil
		go wk.work(wm.workerChan)
	}
}

type worker struct {
	id  int
	err error
}

func (wk *worker) work(workerChan chan<- *worker) {
	var err error
	defer func() {
		if r := recover(); r != nil {
			wk.err = fmt.Errorf("panic happened with [%v]", r)
		} else {
			wk.err = err
		}

		workerChan <- wk
	}()

	fmt.Println("Start Worker...ID = ", wk.id)

	for i := 0; i < 5; i++ {
		time.Sleep(time.Second * 1)
	}

	if rand.Intn(10) > 5 {
		panic("worker panic..")
	} else {
		err = errors.New("work err")
		runtime.Goexit()
	}
}
