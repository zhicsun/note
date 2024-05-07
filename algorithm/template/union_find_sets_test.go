package template

import (
	"testing"
)

func TestUnionFindSets(t *testing.T) {
	tests := []struct {
		name string
		s    [][]int
		want int
	}{
		{"测试1",
			[][]int{
				{1, 1, 0},
				{1, 1, 0},
				{0, 0, 1},
			},
			2,
		},

		{"测试2",
			[][]int{
				{1, 0, 0}, // 00
				{0, 1, 0}, // 11
				{0, 0, 1}, // 21, 22
			},
			3,
		},

		{"测试3",
			[][]int{
				{1, 0, 0, 0},
				{0, 1, 0, 0},
				{0, 0, 1, 0},
				{0, 0, 0, 1},
			},
			4,
		},

		{"测试4",
			[][]int{
				{1, 0, 0, 0},
				{1, 0, 0, 0},
				{0, 0, 1, 0},
				{0, 0, 0, 1},
			},
			3,
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			r := unionFindSets(v.s)
			if r != v.want {
				t.Fatal("结果: ", r, "期望值: ", v.want)
			}
		})
	}
}

func unionFindSets(s [][]int) int { // 二维数组的值表示联通关系
	r := len(s)              //总的数量
	p := make([]int, len(s)) //联通关系

	for i := 0; i < len(p); i++ { // 初始化
		p[i] = -1
	}

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[i]); j++ {
			if s[i][j] == 1 && i != j { // 两个值相连且不是自身
				if union(p, i, j) { // 两个值相连总数相减
					r--
				}
			}
		}
	}

	return r
}

func union(p []int, i, j int) bool {
	u := false       // 两个值是否联通
	pi := find(p, i) // 求相连父值
	pj := find(p, j) // 求相连父值
	if pi != pj {    // 父值不相同进行连接和改变状态
		p[pi] = pj
		u = true
	}

	return u
}

func find(p []int, i int) int { // 递归寻找父值
	if p[i] == -1 { // 父值不存在，返回本身，递归终止条件
		return i
	}

	return find(p, p[i]) // 继续寻找父值
}
