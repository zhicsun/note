package template

import (
	"reflect"
	"testing"
)

func TestMonotonicQueue(t *testing.T) {
	tests := []struct {
		name   string
		arr    []int
		window int
		want   []int
	}{
		{"一个", []int{1}, 1, []int{1}},
		{"多个", []int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{-1, -3, -3, -3, 3, 3}},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if r := MonotonicQueue(v.arr, v.window); !reflect.DeepEqual(r, v.want) {
				t.Fatal("结果: ", r, "期望值: ", v.want)
			}
		})
	}
}

func MonotonicQueue(arr []int, w int) []int {
	r := make([]int, 0)
	queue := make([]int, 0)

	for i := 0; i < len(arr); i++ {
		for len(queue) > 0 && i-w+1 > queue[0] {
			queue = queue[1:]
		}

		for len(queue) > 0 && arr[queue[len(queue)-1]] > arr[i] {
			queue = queue[:len(queue)-1]
		}

		queue = append(queue, i)
		if i >= w-1 {
			r = append(r, arr[queue[0]])
		}
	}

	return r
}
