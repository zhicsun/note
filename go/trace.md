# trace

## 示例

### 脚本

```go
package main

import (
	"os"
	"runtime/trace"
)

func main() {
	trace.Start(os.Stdout)
	defer trace.Stop()

	ch := make(chan string)
	go func() {
		ch <- "trace"
	}()
	<-ch
}

```

### http

```go
package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"sync"
	"time"
)

func init() {
	runtime.SetMutexProfileFraction(1)
	runtime.SetBlockProfileRate(1)
}

var data = make([]string, 0)

func main() {
	go func() {
		for {
			fmt.Println(operate("pprof"))
			time.Sleep(time.Second)
		}
	}()

	var m sync.Mutex
	var data = make(map[int]struct{})
	for i := 0; i < 999; i++ {
		go func(i int) {
			m.Lock()
			defer m.Unlock()
			data[i] = struct{}{}
		}(i)
	}

	err := http.ListenAndServe(":8086", nil)
	if err != nil {
		panic(err)
	}
}

func operate(str string) int {
	data = append(data, str)
	return len(data)
}

```

## 获取 trace 文件

### 脚本

```shell
go run main.go > trace
```

### http

```shell
wget http://127.0.0.1:8086/debug/pprof/trace\?seconds\=6 -o trace
```

## 分析

```shell
go tool trace trace
```

### 首页详情

- View trace：查看跟踪。
- Goroutine analysis：goroutine 分析。
- Network blocking profile：网络阻塞概况。
- Synchronization blocking profile：同步阻塞概况。
- Syscall blocking profile：系统调用阻塞概况。
- Scheduler latency profile：调度延迟概况。
- User defined tasks：用户自定义任务。
- User defined regions：用户自定义区域。
- Minimum mutator utilization：最低 mutator 利用率

### Scheduler latency profile

一般来说，应先查看 Scheduler latency profile 看整体的调用开销情况。

### Goroutine 分析

通过 Goroutine 可以看到在整个运行过程中，每个函数块有多少个 goroutine 在执行，并且每个 Goroutine 的运行开销都花费在哪个阶段。

Goroutine 状态

- Execution Time：执行时间
- Network Wait Time：网络等待时间
- Sync Block Time：同步阻塞时间
- Blocking Syscall Time：调用阻塞时间
- Scheduler Wait Time：调度等待时间
- GC sweeping：GC 清扫
- GC pause：GC 暂停

### View trace 分析

#### 概览

- 时间线：体可按组合键 shift + ？查看帮助手册。
- 堆：显示执行期间的内存分配和释放情况。
- 协程：显示在执行期间每个 goroutine 运行阶段有多少个协程在运行，包含 GC 等待（GCWaiting）、可运行（Runnable）和运行中（Running）三种状态。
- OS 线程：显示在执行期间有多少个线程在运行，包含正在调用 Syscall（InSyscall）和运行中（Running）两种状态。
- 虚拟处理器：每个虚拟处理器显示一行，虚拟处理器的数量一般默认为系统内核数。
- 协程和事件：显示在每个虚拟处理器上有哪些 goroutine 正在运行，而连线行为代表事件关联。

### Goroutine

- Start time：开始时间。
- Wall duration：持续时间。
- Self time：执行时间。
- Start stack trace：开始时的堆栈信息。
- End stack trace：结束时的堆栈信息。
- Incoming flow：输入流。
- Outgoing flow：输出流。
- Preceding events：之前的事件。
- Following events：之后的事件。
- All connected：所有连接的事件。

### Following events

可以通过单击 Following events 等，查看应用运行中的事件流情况。