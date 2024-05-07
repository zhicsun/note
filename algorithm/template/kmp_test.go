package template

import (
	"reflect"
	"testing"
)

func TestKmp(t *testing.T) {
	tests := []struct {
		name, s, p string
		want       []int
	}{
		{"开始", "ababa", "aba", []int{0, 2}},
		{"中间", "aababaa", "aba", []int{1, 3}},
		{"中间", "abcaba", "aba", []int{3}},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if r := kmp(v.s, v.p); !reflect.DeepEqual(r, v.want) {
				t.Fatal("结果: ", r, "期望值: ", v.want)
			}
		})
	}
}

func kmp(s, p string) []int {
	s = " " + s
	p = " " + p

	n := make([]int, len(p))
	for i, j := 2, 0; i < len(p); i++ {
		for j > 0 && p[i] != p[j+1] {
			j = n[j]
		}

		if p[i] == p[j+1] {
			j++
		}

		n[i] = j
	}

	r := make([]int, 0)
	for i, j := 1, 0; i < len(s); i++ {
		for j > 0 && s[i] != p[j+1] {
			j = n[j]
		}

		if s[i] == p[j+1] {
			j++
		}

		if j == len(p)-1 {
			r = append(r, i-j)
			j = n[j]
		}
	}

	return r
}
