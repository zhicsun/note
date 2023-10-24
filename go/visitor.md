# 访问者模式

这个模式是将算法与操作对象的结构分离的一种方法，这种分离的实际结果是能够在不修改结构的情况下向现有对象结构添加新操作，是遵循开放封闭原则的一种方法。

````go
package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

func main() {
	c := Circle{10}
	r := Rectangle{100, 200}
	shapes := []Shape{c, r}

	for _, s := range shapes {
		s.accept(JsonVisitor)
		s.accept(XmlVisitor)
	}

}

type Visitor func(Shape)

type Shape interface {
	accept(Visitor)
}

type Circle struct {
	Radius int
}

func (c Circle) accept(v Visitor) {
	v(c)
}

type Rectangle struct {
	Width, High int
}

func (r Rectangle) accept(v Visitor) {
	v(r)
}

func JsonVisitor(shape Shape) {
	bytes, err := json.Marshal(shape)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func XmlVisitor(shape Shape) {
	bytes, err := xml.Marshal(shape)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}
````