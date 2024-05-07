package template

import (
	"reflect"
	"testing"
)

func TestMonotonicStack(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{"一个", []int{1}, []int{-1}},
		{"多个", []int{3, 4, 2, 7, 8, 5}, []int{-1, 3, -1, 2, 7, 2}},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if r := MonotonicStack(v.arr); !reflect.DeepEqual(r, v.want) {
				t.Fatal("结果: ", r, "期望值: ", v.want)
			}
		})
	}
}

func MonotonicStack(arr []int) []int {
	r := make([]int, 0)
	stack := make([]int, 0)

	for i := 0; i < len(arr); i++ {
		for len(stack) > 0 && arr[stack[len(stack)-1]] > arr[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			r = append(r, -1)
		} else {
			r = append(r, arr[stack[len(stack)-1]])
		}

		stack = append(stack, i)
	}

	return r
}
