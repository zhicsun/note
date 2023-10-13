# 回溯

## [39. 组合总和](https://leetcode.cn/problems/combination-sum/description/)

```go
package main

func combinationSum(candidates []int, target int) [][]int {
	// 初始化返回结果和选择路径
	res := make([][]int, 0)
	path := make([]int, 0)

	// 定义递归函数
	var dfs func(start, sum int, path []int)
	dfs = func(start, sum int, path []int) {
		// 所有路径和大于目标值，结束递归
		if sum > target {
			return
		}

		// 路径和等于目标值，加入到返回结果中
		if sum == target {
			t := make([]int, len(path))
			copy(t, path)
			res = append(res, t)
			return
		}

		// 从当前位置开始遍历
		for i := start; i < len(candidates); i++ {
			// 选择路径
			path = append(path, candidates[i])
			// 所有路径和
			sum += candidates[i]
			dfs(i, sum, path)
			sum -= candidates[i]
			path = path[:len(path)-1]
		}
	}

	dfs(0, 0, path)
	return res
}

```

## [46. 全排列](https://leetcode.cn/problems/permutations/description/)

```go
package main

func permute(nums []int) [][]int {
	// 初始化返回结果
	res := make([][]int, 0)
	// 初始化已选择节点集合
	path := make([]int, 0)
	// 初始化已使用节点集合
	used := make(map[int]bool, 0)

	// 定义递归函数
	var dfs func(nums, path []int, used map[int]bool)
	dfs = func(nums, path []int, used map[int]bool) {
		// 递归终止条件
		if len(nums) == len(path) {
			t := make([]int, len(path))
			copy(t, path)
			res = append(res, t)
			return
		}

		// 遍历所有选择
		for i := 0; i < len(nums); i++ {
			// 跳过已经选择的节点
			if used[i] {
				continue
			}

			// 添加本次选择节点到集合中
			path = append(path, nums[i])
			// 添加本次选择节点到已使用节点集合中
			used[i] = true
			dfs(nums, path, used)
			// 恢复已使用节点集合
			used[i] = false
			// 恢复节点集合
			path = path[0 : len(path)-1]
		}
	}

	// 调用函数
	dfs(nums, path, used)
	// 返回结果
	return res
}

```

## [47. 全排列 II](https://leetcode.cn/problems/permutations-ii/)

```go
package main

import "sort"

func permuteUnique(nums []int) [][]int {
	// 初始化返回结果和已选择节点集合
	res := make([][]int, 0)
	path := make([]int, 0)
	n := len(nums)
	used := make(map[int]bool, n)
	sort.Ints(nums)

	// 定义递归函数
	var dfs func(path []int, used map[int]bool)
	dfs = func(path []int, used map[int]bool) {
		// 递归终止条件
		if len(path) == n {
			t := make([]int, n)
			copy(t, path)
			res = append(res, path)
			return
		}

		// 循环所有节点
		for i := 0; i < n; i++ {
			// 当钱节点已使用跳过
			if used[i] {
				continue
			}

			// 有重复节点
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}

			// 添加本次选择节点到集合中
			path = append(path, nums[i])
			// 添加本次选择节点到已使用节点集合中
			used[i] = true
			dfs(path, used)
			// 恢复已使用节点集合
			used[i] = false
			// 恢复节点集合
			path = path[:len(path)-1]
		}
	}

	dfs(path, used)
	return res
}
```

## [51. N 皇后](https://leetcode.cn/problems/n-queens/description/)

```go
package main

import "strings"

func solveNQueens(n int) [][]string {
	// 初始化返回结果
	res := make([][]string, 0)
	// 初始化棋盘
	board := make([]string, n)
	for i := 0; i < n; i++ {
		board[i] = strings.Repeat(".", n)
	}

	// 定义递归函数
	var dfs func(board []string, row int)
	dfs = func(board []string, row int) {
		// 终止条件，到棋盘最后一行
		if row == n {
			t := make([]string, n)
			copy(t, board)
			res = append(res, t)
			return
		}

		// 循环列举每一列
		for col := 0; col < n; col++ {
			// 过滤不符合的
			if !isValid(board, row, col) {
				continue
			}

			// 进行回溯
			line := []byte(board[row])
			line[col] = 'Q'
			board[row] = string(line)
			dfs(board, row+1)
			line[col] = '.'
			board[row] = string(line)
		}
	}

	// 调用函数
	dfs(board, 0)
	// 返回结果
	return res
}

func isValid(board []string, row, col int) bool {
	n := len(board)
	// 检查列是否有皇后冲突
	for i := 0; i < n; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}
	// 检查右上方是否有皇后冲突
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == 'Q' {
			return false
		}
	}
	// 检查左上方是否有皇后冲突
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}
	return true
}

```

## [77. 组合](https://leetcode.cn/problems/combinations/description/)

```go
package main

func combine(n int, k int) [][]int {
	// 初始化返回结果和选择的路径
	res := make([][]int, 0)
	path := make([]int, 0)

	// 定义和实现递归函数
	var dfs func(idx int, path []int)
	dfs = func(idx int, path []int) {
		// 选择路径的数量和要求的数量相同，递归终止
		if len(path) == k {
			t := make([]int, k)
			copy(t, path)
			res = append(res, t)
			return
		}

		// 循环剩余的子集位置
		for start := idx; start <= n; start++ {
			path = append(path, start)
			dfs(start+1, path)
			path = path[:len(path)-1]
		}
	}

	dfs(1, path)
	return res
}

```

## [78. 子集](https://leetcode.cn/problems/subsets/description/)

```go
package main

func subsets(nums []int) [][]int {
	// 初始化返回结果和子集
	res := make([][]int, 0)
	path := make([]int, 0)

	// 定义递归函数
	var dfs func(start int, paht []int)

	// 实现递归函数
	dfs = func(start int, path []int) {
		// 把子集加入返回结果
		t := make([]int, len(path))
		copy(t, path)
		res = append(res, t)

		// 循环剩余的子集位置
		for i := start; i < len(nums); i++ {
			// 添加当前使用的节点
			path = append(path, nums[i])
			dfs(i+1, path)
			path = path[:len(path)-1]
		}
	}

	dfs(0, path)
	return res
}

```

## [90. 子集 II](https://leetcode.cn/problems/subsets-ii/)

```go
package main

import "sort"

func subsetsWithDup(nums []int) [][]int {
	// 初始化返回结果和选择路径
	res := make([][]int, 0)
	path := make([]int, 0)

	// 排序便于后面去重
	sort.Ints(nums)

	// 定义和实现递归函数
	var dfs func(start int, path []int)
	dfs = func(start int, path []int) {
		// 存储选择的路径
		t := make([]int, len(path))
		copy(t, path)
		res = append(res, t)

		// 循环剩下的节点
		for i := start; i < len(nums); i++ {
			// 有重复跳过
			if i > start && nums[i] == nums[i-1] {
				continue
			}

			path = append(path, nums[i])
			dfs(i+1, path)
			path = path[:len(path)-1]
		}
	}

	dfs(0, path)
	return res
}

```
