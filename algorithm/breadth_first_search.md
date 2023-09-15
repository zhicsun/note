# 广度优先搜索

## [111. 二叉树的最小深度](https://leetcode.cn/problems/minimum-depth-of-binary-tree/)

```go
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	// 节点为空直接返回
	if root == nil {
		return 0
	}

	// 初始化深度和队列
	depth := 1
	q := make([]*TreeNode, 0)
	q = append(q, root)

	// 队列有值
	for len(q) > 0 {
		// 计算当前层数量
		l := len(q)
		// 遍历当前层元素
		for i := 0; i < l; i++ {
			// 取出头部节点
			c := q[0]
			q = q[1:]

			// 当前元素师最近的节点，直接返回
			if c.Left == nil && c.Right == nil {
				return depth
			}

			// 入队列下一层元素
			if c.Left != nil {
				q = append(q, c.Left)
			}
			if c.Right != nil {
				q = append(q, c.Right)
			}
		}
		// 进入下一层
		depth++
	}

	// 返回结果
	return depth
}

```

## [752. 打开转盘锁](https://leetcode.cn/problems/open-the-lock/description/)

```go
package main

func openLock(deadends []string, target string) int {
	deads := make(map[string]bool)
	for _, s := range deadends {
		deads[s] = true
	}
	// 用 map 不用 queue，可以快速判断元素是否存在
	q1 := make(map[string]bool)
	q2 := make(map[string]bool)
	visited := make(map[string]bool)

	step := 0
	q1["0000"] = true
	q2[target] = true

	for len(q1) != 0 && len(q2) != 0 {
		// 哈希集合在遍历的过程中不能修改，用 temp 存储扩散结果
		temp := make(map[string]bool)

		/* 将 q1 中的所有节点向周围扩散 */
		for cur := range q1 {
			/* 判断是否到达终点 */
			if _, ok := deads[cur]; ok {
				continue
			}
			if _, ok := q2[cur]; ok {
				return step
			}
			visited[cur] = true

			/* 将一个节点的未遍历相邻节点加入集合 */
			for j := 0; j < 4; j++ {
				up := plusOne(cur, j)
				if _, ok := visited[up]; !ok {
					temp[up] = true
				}
				down := minusOne(cur, j)
				if _, ok := visited[down]; !ok {
					temp[down] = true
				}
			}
		}
		/* 在这里增加步数 */
		step++
		// temp 相当于 q1
		// 这里交换 q1 q2，下一轮 while 就是扩散 q2
		q1 = q2
		q2 = temp
	}
	return -1
}

func plusOne(s string, j int) string {
	sBytes := []byte(s)
	if sBytes[j] == '9' {
		sBytes[j] = '0'
	} else {
		sBytes[j] += 1
	}
	return string(sBytes)
}

func minusOne(s string, j int) string {
	sBytes := []byte(s)
	if sBytes[j] == '0' {
		sBytes[j] = '9'
	} else {
		sBytes[j] -= 1
	}
	return string(sBytes)
}

```
