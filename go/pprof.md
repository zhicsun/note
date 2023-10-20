# pprof

## 作用

- CPU 分析
- 内存分析
- 阻塞分析
- 互斥锁分析
- Goroutine 分析
- 操作系统线程分析

## 采集方式

- runtime/pprof：采集程序（非 Server）的指定区块的运行数据进行分析。
- net/http/pprof：基于 HTTP Server 运行，并且可以采集运行时数据进行分析。
- go test：通过运行测试用例，并指定所需标识来进行采集。

## 使用方式

- 报告查看。
- 交互式终端。
- Web 界面。

## HTTP Server

代码

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

### 生成报告

url 不带 debug 参数，将会直接下载对应的 profile 文件。

[pprof：报告入口](http://127.0.0.1:8086/debug/pprof/)

[allocs：查看过去所有内存分配的样本](http://127.0.0.1:8086/debug/pprof/allocs?debug=1)

[block：查看导致阻塞同步的堆栈跟踪](http://127.0.0.1:8086/debug/pprof/block?debug=1)

[cmdline： 当前程序的命令行的完整调用路径](http://127.0.0.1:8086/debug/pprof/cmdline?debug=1)

[goroutine：当前所有运行的 goroutines 堆栈跟踪](http://127.0.0.1:8086/debug/pprof/goroutine?debug=1)

[heap：查看活动对象的内存分配情况](http://127.0.0.1:8086/debug/pprof/heap?debug=1)

[mutex：查看导致互斥锁的竞争持有者的堆栈跟踪](http://127.0.0.1:8086/debug/pprof/mutex?debug=1)

[profile： 默认 30s 的 CPU Profiling](http://127.0.0.1:8086/debug/pprof/profile?debug=1)

[threadcreate：查看创建新 OS 线程的堆栈跟踪](http://127.0.0.1:8086/debug/pprof/threadcreate?debug=1)

[trace：当前程序的执行详情](http://127.0.0.1:8086/debug/pprof/trace?debug=1)

### 终端使用

CPU Profiling

```shell
go tool pprof http://localhost:8086/debug/pprof/profile\?seconds\=6

go tool pprof https+insecure://localhost:8086/debug/pprof/profile\?seconds\=6

pprof help

top 10
```

top 输出

- flat：函数自身的运行耗时。
- flat%：函数自身在 CPU 运行耗时总比例。
- sum%：函数自身累积使用 CPU 总比例。
- cum：函数自身及其调用函数的运行总耗时。
- cum%：函数自身及其调用函数的运行耗时总比例。
- Name：函数名。

Heap Profiling

```shell
go tool pprof -inuse_space http://localhost:8086/debug/pprof/heap\?seconds\=6
go tool pprof -alloc_objects http://localhost:8086/debug/pprof/heap\?seconds\=6
go tool pprof -inuse_objects http://localhost:8086/debug/pprof/heap\?seconds\=6
go tool pprof -alloc_space http://localhost:8086/debug/pprof/heap\?seconds\=6

top 10
```

top 输出

- inuse_space：分析应用程序的常驻内存占用情况。
- alloc_space：分析应用程序分配的内存空间大小。
- inuse_objects：分析应用程序的每个函数所分配的对象数量。
- alloc_objects：分析应用程序的内存临时分配情况。

Goroutine Profiling

```shell
go tool pprof http://localhost:8086/debug/pprof/goroutine\?seconds\=6

traces
```

traces

- 打印出对应的所有调用栈。
- 在 Heap Profiling 展示的是占用内存大小。

Mutex Profiling

```shell
go tool pprof http://localhost:8086/debug/pprof/mutex\?seconds\=6
```

top 查看互斥量的排名，list 查看指定函数的代码情况。

Block Profiling

```shell
go tool pprof http://localhost:8086/debug/pprof/block\?seconds\=6
```

top 查看互斥量的排名，list 查看指定函数的代码情况。

### 可视化界面

```shell
brew install graphviz

wget http://127.0.0.1:8086/debug/pprof/profile\?seconds\=6

go tool pprof -http=:8088 profile

go tool pprof profile 
```

top

与命令行一致

Graph

为整体的函数调用流程，框越大、线越粗、框颜色越鲜艳（红色）就代表它占用的时间越久，开销越大。

Peek

相较于 Top 视图，增加了所属的上下文信息的展示，也就是函数的输出调用者/被调用者。

Source

该视图主要是增加了面向源代码的追踪和分析，可以看到其开销主要消耗在哪里。

Flame Graph

Flame Graph（火焰图）它是可动态的，调用顺序由上到下（A -> B -> C -> D），每一块代表一个函数、颜色越鲜艳（红）、区块越大代表占用
CPU 的时间更长。同时它也支持点击块深入进行分析。

## 测试

代码

```go
package main

import "testing"

func TestAdd(t *testing.T) {
	_ = operate("go-programming-tour-book")
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		operate("go-programming-tour-book")
	}
}

```

CPU

```shell
go test -bench=. -cpuprofile=cpu.profile

go tool pprof cpu.profile

web
```

Memory

```shell
go test -bench=. -memprofile=mem.profile

go tool pprof mem.profile

web
```

## goroutine 增多问题排查

多次拉取对比

````shell
go tool pprof http://localhost:8086/debug/pprof/goroutine
go tool pprof http://localhost:8086/debug/pprof/goroutine
go tool pprof -base 文件1 文件2
````

top 看到引起问题的函数

```shell
top
```

traces 得到具体的调用栈

```shell
traces
```

list 查看查看函数详情

```shell
list 异常函数
```