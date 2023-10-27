# 访问者模式

访问者模式可以给一系列对象透明的添加功能，并且把相关代码封装到一个类中，对象只要预留访问者接口Accept则后期为对象添加功能的时候就不需要改动对象。

````go
package main

import (
	"fmt"
	"strings"
)

func main() {
	ming := &PhoneFans{name: "小明"}
	wang := &ComputerFans{name: "小王"}

	shop := Shop{}
	shop.Register(ming, wang)

	shop.SetProduct("手机")
	shop.SetProduct("电脑")
}

type Shop struct {
	Product string
	Buyers  []Buyer
}

func (r *Shop) Register(buyers ...Buyer) {
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
		r.Buyers[i].Inform(*r)
	}
}

type Buyer interface {
	Inform(Shop)
}

type PhoneFans struct {
	name string
}

func (r *PhoneFans) Inform(shop Shop) {
	if strings.Contains(shop.Product, "手机") {
		fmt.Println(r.name, "购买了", shop.Product)
	}
}

type ComputerFans struct {
	name string
}

func (r *ComputerFans) Inform(shop Shop) {
	if strings.Contains(shop.Product, "电脑") {
		fmt.Println(r.name, "购买了", shop.Product)
	}
}

````