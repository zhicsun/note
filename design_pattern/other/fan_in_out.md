# 扇入扇出

扇入和扇出是并发编程模式，用于解决多个输入和输出之间的数据流处理问题，这两个模式通常用于构建高性能、高并发的系统.

```go
package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	nums := makeRange(1, 10000)
	in := echo(nums)

	const nProcess = 5
	var ch [nProcess]<-chan int
	for i := range ch {
		ch[i] = sum(prime(in))
	}

	for n := range sum(merge(ch[:])) {
		fmt.Println(n)
	}
}

func merge(cs []<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	wg.Add(len(cs))
	for _, c := range cs {
		go func(c <-chan int) {
			for n := range c {
				out <- n
			}
			wg.Done()
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func isPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func prime(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			if isPrime(n) {
				out <- n
			}
		}
		close(out)
	}()
	return out
}

func sum(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var sum = 0
		for n := range in {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func echo(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

```