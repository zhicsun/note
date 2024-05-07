package paradigm

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"testing"
)

func TestChannelSync(t *testing.T) {
	limit := NewChannelSync(10)
	limit.Start(math.MaxInt16)
	limit.wg.Wait()
}

type ChannelSync struct {
	wg    sync.WaitGroup
	ch    chan struct{}
	limit int
}

func NewChannelSync(limit int) *ChannelSync {
	return &ChannelSync{
		ch:    make(chan struct{}, limit),
		limit: limit,
	}
}

func (r *ChannelSync) Start(n int) {
	for i := 0; i < n; i++ {
		r.wg.Add(1)
		r.ch <- struct{}{}
		go r.deal(i)
	}
}

func (r *ChannelSync) deal(i int) {
	defer func() {
		<-r.ch
		r.wg.Done()
	}()
	fmt.Println(i, runtime.NumGoroutine())
}
