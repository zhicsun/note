package paradigm

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"testing"
)

func TestChannelSyncDetach(t *testing.T) {
	limit := NewChannelSyncDetach(10)
	limit.Deal()
	limit.Send(math.MaxInt8)
	limit.wg.Wait()
}

type ChannelSyncDetach struct {
	wg    sync.WaitGroup
	ch    chan int
	limit int
}

func NewChannelSyncDetach(limit int) *ChannelSyncDetach {
	return &ChannelSyncDetach{
		ch:    make(chan int, limit),
		limit: limit,
	}
}

func (r *ChannelSyncDetach) Deal() {
	for i := 0; i < r.limit; i++ {
		go func(i int) {
			for v := range r.ch {
				r.wg.Done()
				fmt.Println(v, i, runtime.NumGoroutine())
			}
		}(i)
	}
}

func (r *ChannelSyncDetach) Send(n int) {
	for i := 0; i < n; i++ {
		r.wg.Add(1)
		r.ch <- i
	}
}
