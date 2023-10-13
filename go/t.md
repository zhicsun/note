```go
package main

import (
	"fmt"
	"time"
)

func main() {
	total := 0
	count := 30
	op := NewOperator()

	for i := 0; i < count; i++ {
		total += op.Normal(i)
	}
	println("total time cost:", total)

	concurrentNumber := 3
	data := make(chan int)
	for i := 0; i < count; i++ {
		data <- i
	}
	fmt.Println(op.Concurrent(concurrentNumber, data))
}

func NewOperator() Operator {
	return Operator{}
}

type Operator struct{}

func (o Operator) Concurrent(cNumber int, data <-chan int) []chan int {
	result := make([]chan int, cNumber)
	for i := 0; i < cNumber; i++ {
		go func(i int) {
			t := 0
			for {
				_, ok := <-data
				if !ok {
					result[i] <- t
				}
				t += o.Normal(i)
			}
		}(i)
	}
	return result
}

func (o Operator) Normal(id int) int {
	total := 0
	println("goroutine-", id, ": exec ...")
	total += o.fn1(id)
	total += o.fn2(id)
	println("goroutine-", id, ": exec done")
	return total
}

func (Operator) fn1(id int) int {
	cost := 60
	time.Sleep(time.Millisecond * time.Duration(cost))
	println("goroutine-", id, ": fn1 ok")
	return cost
}

func (Operator) fn2(id int) int {
	cost := 120
	time.Sleep(time.Millisecond * time.Duration(cost))
	println("goroutine-", id, ": fn2 ok")
	return cost
}

```