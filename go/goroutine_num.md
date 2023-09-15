# goroutine 数量控制

## 不控制 goroutine 数量引发的问题

- cpu 和 memory 飙升
- 主进程崩溃

## 控制 goroutine 数量的方法

### buffer channel

```go
package main

import (
	"fmt"
	"math"
	"runtime"
)

var buffer = make(chan struct{}, 10)
var count = math.MaxInt

func main() {
	for i := 0; i < count; i++ {
		buffer <- struct{}{}
		go deal(i)
	}
}

func deal(i int) {
	fmt.Println(i, runtime.NumGoroutine())
	<-buffer
}

```

如果 for 循环次数太少，会导致在当主协程结束时，子协程也是会被终止掉来不及把值输出。

### channel 和 sync 组合

```go
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var buffer = make(chan struct{}, 10)
var count = 1

func main() {
	for i := 0; i < count; i++ {
		wg.Add(1)
		buffer <- struct{}{}
		go deal(i)
	}
	wg.Wait()
}

func deal(i int) {
	defer wg.Done()
	fmt.Println(i, runtime.NumGoroutine())
	<-buffer
}

```

### 发送和接收分离

```go
package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var buffer = make(chan int)
var count = math.MaxInt64
var goroutineCount = 10

func main() {
	for i := 0; i < goroutineCount; i++ {
		go deal(i)
	}

	for i := 0; i < count; i++ {
		send(i)
	}

	wg.Wait()
}

func send(i int) {
	wg.Add(1)
	buffer <- i
}

func deal(i int) {
	defer wg.Done()
	for v := range buffer {
		fmt.Println(v, runtime.NumGoroutine())
	}
}

```