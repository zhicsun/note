# 多函数并发

```go
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	count := 30
	queue := make(chan int)
	gNumber := 3
	wg := &sync.WaitGroup{}
	for i := 0; i < gNumber; i++ {
		exec(wg, i+1, queue)
	}

	for i := 0; i < count; i++ {
		wg.Add(1)
		queue <- i
	}
	wg.Wait()
	close(queue)
	time.Sleep(1000 * time.Second)
}

func fn1(id, data int) int {
	delay := 60
	time.Sleep(time.Millisecond * time.Duration(delay))
	fmt.Println(id, data, runtime.NumGoroutine())
	return delay
}

func fn2(id, data int) int {
	delay := 120
	time.Sleep(time.Millisecond * time.Duration(delay))
	fmt.Println(id, data, runtime.NumGoroutine())
	return delay
}

func fn3(id, data int) int {
	delay := 180
	time.Sleep(time.Millisecond * time.Duration(delay))
	fmt.Println(id, data, runtime.NumGoroutine())
	return delay
}

func start(wg *sync.WaitGroup, id int, f func(int, int) int, next chan<- int) (chan<- int, chan<- struct{}, <-chan int) {
	queue := make(chan int)
	quit := make(chan struct{})
	result := make(chan int)

	go func() {
		total := 0
		for {
			select {
			case <-quit:
				result <- total
				return
			case v := <-queue:
				total += f(id, v)
				if next != nil {
					next <- v
				} else {
					wg.Done()
				}
			}
		}

	}()
	return queue, quit, result
}

func exec(wg *sync.WaitGroup, id int, queue <-chan int) {
	go func(id int) {
		queue3, quit3, result3 := start(wg, id, fn3, nil)
		queue2, quit2, result2 := start(wg, id, fn2, queue3)
		queue1, quit1, result1 := start(wg, id, fn1, queue2)

		for {
			select {
			case v, ok := <-queue:
				if !ok {
					close(quit1)
					close(quit2)
					close(quit3)
					total := max(<-result1, <-result2, <-result3)
					fmt.Println("goroutine", id, "time cost:", total)
					return
				}
				queue1 <- v
			}
		}
	}(id)
}

func max(args ...int) int {
	n := 0
	for i := 0; i < len(args); i++ {
		if args[i] > n {
			n = args[i]
		}
	}
	return n
}

```