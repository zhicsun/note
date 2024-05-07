package paradigm

import (
	"fmt"
	"golang.org/x/time/rate"
	"sync"
	"testing"
	"time"
)

var wg sync.WaitGroup

func TestNormal(t *testing.T) {
	limiter := rate.NewLimiter(rate.Limit(200), 100) // 第一个参数每秒产生的令牌，第二个参数是令牌桶的初始容量
	for i := 0; i < 200; i++ {
		if !limiter.Allow() {
			fmt.Println("Rate limit exceeded. Request rejected.")
			continue
		}
		wg.Add(1)
		go process()
	}
	wg.Wait()
}

func process() {
	wg.Done()
	fmt.Println("Request processed successfully.")
	time.Sleep(time.Millisecond)
}
