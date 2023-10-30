# 访问者模式

主要解决的是数据与算法的耦合问题，尤其是在数据结构比较稳定，而算法多变的情况下。为了不“污染”数据本身，访问者模式会将多种算法独立归类，并在访问数据时根据数据类型自动切换到对应的算法，实现数据的自动响应机制，并且确保算法的自由扩展。

```go
package main

import "fmt"

func main() {
	container := []Container{
		&CandyContainer{
			ProductContainer{Price: 20.1},
		},
		&WineContainer{
			ProductContainer{Price: 30.2},
		},
		&FruitContainer{
			ProductContainer: ProductContainer{Price: 10.0},
			weight:           86.86,
		},
	}

	discountVisitor := DiscountVisitor{}
	for i := 0; i < len(container); i++ {
		discountVisitor.Visit(container[i])
	}
}

type Container interface {
	Accept(Visitor)
	GetPrice() float64
}

type Visitor interface {
	Visit(Container)
}

type ProductContainer struct {
	Price float64
}

func (r *ProductContainer) SetPrice(price float64) {
	r.Price = price
}

func (r *ProductContainer) GetPrice() float64 {
	return r.Price
}

type CandyContainer struct {
	ProductContainer
}

func (r *CandyContainer) Accept(visitor Visitor) {
	visitor.Visit(r)
}

type WineContainer struct {
	ProductContainer
}

func (r *WineContainer) Accept(visitor Visitor) {
	visitor.Visit(r)
}

type FruitContainer struct {
	ProductContainer
	weight float64
}

func (r *FruitContainer) SetWeight(weight float64) {
	r.weight = weight
}

func (r *FruitContainer) GetWeight() float64 {
	return r.weight
}

func (r *FruitContainer) Accept(visitor Visitor) {
	visitor.Visit(r)
}

type CandyVisitor struct{}

func (r *CandyVisitor) Visit(container Container) {
	fmt.Println("糖果价格", container.GetPrice())
}

type WineVisitor struct{}

func (r *WineVisitor) Visit(container Container) {
	fmt.Println("酒价格", container.GetPrice()*6)
}

type FruitVisitor struct{}

func (r *FruitVisitor) Visit(container Container) {
	fmt.Println("水果价格", container.GetPrice()*8)
}

type DiscountVisitor struct{}

func (r *DiscountVisitor) Visit(container Container) {
	switch container.(type) {
	case *CandyContainer:
		visitor := &CandyVisitor{}
		container.Accept(visitor)
	case *WineContainer:
		visitor := &WineVisitor{}
		container.Accept(visitor)
	case *FruitContainer:
		visitor := &FruitVisitor{}
		container.Accept(visitor)
	}
}
```