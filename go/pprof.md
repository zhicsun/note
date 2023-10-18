# pprof

## 有什么用

CPU 分析

内存分析

阻塞分析

互斥锁分析

Goroutine 分析

操作系统线程分析

## 采集方式

runtime/pprof：采集程序（非 Server）的指定区块的运行数据进行分析。

net/http/pprof：基于 HTTP Server 运行，并且可以采集运行时数据进行分析。

go test：通过运行测试用例，并指定所需标识来进行采集。

## 使用方式

报告生成。

交互式终端使用。

Web 界面。

## HTTP Server

### 示例代码

```go
package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

var data = make([]string, 0)

func main() {
	go func() {
		for {
			fmt.Println(operate("pprof"))
		}
	}()

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

### 报告查看

[报告概览地址](http://127.0.0.1:8086/debug/pprof/)

相关内容：

[allocs：查看过去所有内存分配的样本](http://127.0.0.1:8086/debug/pprof/allocs?debug=1)

[block：查看导致阻塞同步的堆栈跟踪](http://127.0.0.1:8086/debug/pprof/block?debug=1)

[cmdline： 当前程序的命令行的完整调用路径](http://127.0.0.1:8086/debug/pprof/cmdline?debug=1)

[goroutine：当前所有运行的 goroutines 堆栈跟踪](http://127.0.0.1:8086/debug/pprof/goroutine?debug=1)

[heap：查看活动对象的内存分配情况](http://127.0.0.1:8086/debug/pprof/heap?debug=1)

[mutex：查看导致互斥锁的竞争持有者的堆栈跟踪](http://127.0.0.1:8086/debug/pprof/mutex?debug=1)

[profile： 默认 30s 的 CPU Profiling](http://127.0.0.1:8086/debug/pprof/profile?debug=1)

[threadcreate：查看创建新 OS 线程的堆栈跟踪](http://127.0.0.1:8086/debug/pprof/threadcreate?debug=1)

[trace：当前程序的执行详情](http://127.0.0.1:8086/debug/pprof/trace?debug=1)
