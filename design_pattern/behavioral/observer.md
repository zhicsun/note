# 观察者模式

观察者模式可以针对被观察对象与观察者对象之间一对多的依赖关系建立起一种行为自动触发机制，当被观察对象状态发生变化时主动对外发起广播，以通知所有观察者做出响应。

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	ming := &PhoneFansObserver{name: "小明"}
	wang := &ComputerFansObserver{name: "小王"}

	shop := Shop{}
	shop.Register(ming, wang)

	shop.SetProduct("手机")
	shop.SetProduct("电脑")
}

type Subject interface {
	Register(buyers ...Observer)
	Notify()
}

type Observer interface {
	Update(shop Shop)
}

type Shop struct {
	Product string
	Buyers  []Observer
}

func (r *Shop) Register(buyers ...Observer) {
	for i := 0; i < len(buyers); i++ {
		r.Buyers = append(r.Buyers, buyers[i])
	}
}

func (r *Shop) GetProduct() string {
	return r.Product
}

func (r *Shop) SetProduct(product string) {
	r.Product = product
	r.Notify()
}

func (r *Shop) Notify() {
	for i := 0; i < len(r.Buyers); i++ {
		r.Buyers[i].Update(*r)
	}
}

type PhoneFansObserver struct {
	name string
}

func (r *PhoneFansObserver) Update(shop Shop) {
	if strings.Contains(shop.Product, "手机") {
		fmt.Println(r.name, "购买了", shop.Product)
	}
}

type ComputerFansObserver struct {
	name string
}

func (r *ComputerFansObserver) Update(shop Shop) {
	if strings.Contains(shop.Product, "电脑") {
		fmt.Println(r.name, "购买了", shop.Product)
	}
}

```