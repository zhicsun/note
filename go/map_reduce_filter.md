# Map-Reduce-Filter

## Map 示例

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	var list = []string{"Abc", "Def", "Ghi"}

	x := MapStrToStr(list, func(s string) string {
		return strings.ToUpper(s)
	})
	fmt.Println(x)

	y := MapStrToInt(list, func(s string) int {
		return len(s)
	})
	fmt.Println(y)
}

func MapStrToStr(arr []string, fn func(s string) string) []string {
	var newArray = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		newArray = append(newArray, fn(arr[i]))
	}
	return newArray
}

func MapStrToInt(arr []string, fn func(s string) int) []int {
	var newArray = make([]int, 0)
	for i := 0; i < len(arr); i++ {
		newArray = append(newArray, fn(arr[i]))
	}
	return newArray
}

```

## Reduce 示例

```go
package main

import (
	"fmt"
)

func main() {
	var list = []string{"Abc", "Def", "Ghi"}
	x := Reduce(list, func(s string) int {
		return len(s)
	})

	fmt.Printf("%v\n", x)
}

func Reduce(arr []string, fn func(s string) int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += fn(arr[i])
	}

	return sum
}

```

## 完整示例

```go
package main

import "fmt"

func main() {
	list := []Employee{
		{"Hao", 44, 0, 8000},
		{"Bob", 34, 10, 5000},
		{"Alice", 23, 5, 9000},
		{"Jack", 26, 0, 4000},
		{"Tom", 48, 9, 7500},
		{"Marry", 29, 0, 6000},
		{"Mike", 32, 8, 4000},
	}

	old := EmployeeCountIf(list, func(e *Employee) bool {
		return e.Age > 40
	})
	fmt.Println(old)

	highPay := EmployeeCountIf(list, func(e *Employee) bool {
		return e.Salary > 6000
	})
	fmt.Println(highPay)

	noVacation := EmployeeFilterIn(list, func(e *Employee) bool {
		return e.Vacation == 0
	})
	fmt.Println(noVacation)

	totalPay := EmployeeSumIf(list, func(e *Employee) int {
		return e.Salary
	})
	fmt.Println(totalPay)

	youngerPay := EmployeeSumIf(list, func(e *Employee) int {
		if e.Age < 30 {
			return e.Salary
		}
		return 0
	})
	fmt.Println(youngerPay)
}

type Employee struct {
	Name     string
	Age      int
	Vacation int
	Salary   int
}

func EmployeeCountIf(list []Employee, fn func(e *Employee) bool) int {
	count := 0
	for i := 0; i < len(list); i++ {
		if fn(&list[i]) {
			count += 1
		}
	}
	return count
}

func EmployeeFilterIn(list []Employee, fn func(e *Employee) bool) []Employee {
	var newList []Employee
	for i := 0; i < len(list); i++ {
		if fn(&list[i]) {
			newList = append(newList, list[i])
		}
	}
	return newList
}

func EmployeeSumIf(list []Employee, fn func(e *Employee) int) int {
	var sum = 0
	for i := 0; i < len(list); i++ {
		sum += fn(&list[i])
	}
	return sum
}

```