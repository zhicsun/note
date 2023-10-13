# 控制 goroutine 数量的方法

## channel 和 sync 组合

```go
package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	goroutineCount := 10
	goroutineCountBuffer := make(chan struct{}, goroutineCount)
	count := math.MaxInt64
	for i := 0; i < count; i++ {
		wg.Add(1)
		goroutineCountBuffer <- struct{}{}
		go deal(&wg, goroutineCountBuffer, i)
	}
	wg.Wait()
}

func deal(wg *sync.WaitGroup, buffer <-chan struct{}, i int) {
	defer func() {
		wg.Done()
		<-buffer
	}()
	fmt.Println(i, runtime.NumGoroutine())
}

```

## 发送和接收分离

```go
package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	buffer := make(chan int)
	count := math.MaxInt64
	goroutineCount := 10

	for i := 0; i < goroutineCount; i++ {
		go deal(&wg, buffer, i)
	}

	for i := 0; i < count; i++ {
		send(&wg, i, buffer)
	}

	wg.Wait()
}

func send(wg *sync.WaitGroup, i int, buffer chan<- int) {
	wg.Add(1)
	buffer <- i
}

func deal(wg *sync.WaitGroup, buffer <-chan int, i int) {
	defer wg.Done()
	for v := range buffer {
		fmt.Println(v, i, runtime.NumGoroutine())
	}
}

```