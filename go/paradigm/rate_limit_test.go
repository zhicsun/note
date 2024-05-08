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
	limiter := rate.NewLimiter(rate.Limit(10), 190) // 第一个参数每秒产生的令牌，第二个参数是令牌桶的初始容量
	for i := 0; i < 200; i++ {
		if !limiter.Allow() {
			limiter.SetLimit(rate.Limit(300))
			fmt.Println("Rate limit exceeded. Request rejected.")
			time.Sleep(time.Second)
			continue
		}
		go process()
	}
	wg.Wait()
}

func process() {
	defer wg.Done()
	wg.Add(1)
	fmt.Println("Request processed successfully.")
	time.Sleep(time.Millisecond)
}
