package paradigm

import (
	"fmt"
	"strings"
	"testing"
)

func TestMapReduce(t *testing.T) {
	list := []string{"Abc", "Def", "Ghi"}
	x := MapStrToStr(list, func(s string) string {
		return strings.ToUpper(s)
	})
	fmt.Println(x)

	y := Reduce(list, func(s string) int {
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

func Reduce(arr []string, fn func(s string) int) int {
	s := 0
	for i := 0; i < len(arr); i++ {
		s += fn(arr[i])
	}

	return s
}
