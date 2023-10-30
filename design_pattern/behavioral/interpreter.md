# 解释器模式

针对某种语言并基于其语法特征创建一系列的表达式类（包括终极表达式与非终极表达式），利用树结构模式将表达式对象组装起来，最终将其翻译成计算机能够识别并执行的语义树。

```go
package main

import "fmt"

func main() {
	sequence := Sequence{expressions: []Expression{
		&Move{
			x: 60,
			y: 80,
		},
		&Repetition{
			LoopCount: 5,
			LoopBodySequence: &Sequence{expressions: []Expression{
				&LeftClick{
					LeftKeyUp:   LeftKeyUp{},
					LeftKeyDown: LeftKeyDown{},
				},
				&Delay{Seconds: 6},
			}},
		},
		&RightKeyDown{},
		&Delay{Seconds: 80},
	}}
	sequence.Interpret()
}

type Expression interface {
	Interpret()
}

type Move struct {
	x, y int
}

func (r *Move) Interpret() {
	fmt.Println("鼠标移动到 X:", r.x, " Y:", r.y)
}

type LeftKeyDown struct{}

func (r *LeftKeyDown) Interpret() {
	fmt.Println("按下鼠标左键")
}

type LeftKeyUp struct{}

func (r *LeftKeyUp) Interpret() {
	fmt.Println("松开鼠标左键")
}

type Delay struct {
	Seconds int
}

func (r *Delay) Interpret() {
	fmt.Println("系统延迟", r.Seconds, "秒")
}

type LeftClick struct {
	LeftKeyUp
	LeftKeyDown
}

func (r *LeftClick) Interpret() {
	r.LeftKeyDown.Interpret()
	r.LeftKeyUp.Interpret()
}

type RightKeyDown struct{}

func (r *RightKeyDown) Interpret() {
	fmt.Println("按下鼠标右键")
}

type Repetition struct {
	LoopCount        int
	LoopBodySequence Expression
}

func (r *Repetition) Interpret() {
	for i := 0; i < r.LoopCount; i++ {
		r.LoopBodySequence.Interpret()
	}
}

type Sequence struct {
	expressions []Expression
}

func (r *Sequence) Interpret() {
	for i := 0; i < len(r.expressions); i++ {
		r.expressions[i].Interpret()
	}
}

```