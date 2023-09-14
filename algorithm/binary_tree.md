# 遍历
## [104. 二叉树的最大深度](https://leetcode.cn/problems/maximum-depth-of-binary-tree/description/)
```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    // 定义返回结果
    res := 0

    // 定义树遍历递归函数
    var dfs func(root *TreeNode, level int)

    // 实现树遍历递归函数
    dfs = func(root *TreeNode, level int) {
        // 节点为空，递归条件终止
        if root == nil {
            return
        }

        // 当前层级大于当前结果，更新当前结果
        if level > res {
            res = level
        }

        // 遍历左子树
        dfs(root.Left, level+1)
        // 编译右子树
        dfs(root.Right, level+1)
    }

    // 调用递归函数
    dfs(root, 1)

    // 返回结果
    return res
}
```
# 分解
## [104. 二叉树的最大深度](https://leetcode.cn/problems/maximum-depth-of-binary-tree/description/)
```go
func maxDepth(root *TreeNode) int {
    // 当前节点为空，递归终止条件
    if root == nil {
        return 0
    }

    // 返回左子树的深度
    left := maxDepth(root.Left)
    // 返回右子树的深度
    right := maxDepth(root.Right)

    // 返回左右子树最大深度加一
    if left > right {
        return left + 1
    } else {
        return right + 1
    }
}
```
## [543. 二叉树的直径](https://leetcode.cn/problems/diameter-of-binary-tree/description/)
```go
func diameterOfBinaryTree(root *TreeNode) int {
    res := 0

    // 定义递归函数
    var dfs func(root *TreeNode) int 
    // 实现递归函数
    dfs = func(root *TreeNode) int {
        // 节点为空，递归终止条件
        if root == nil {
            return 0
        }

        // 获取右子树最大节点数
        l := dfs(root.Left)
        // 获取左子树最大节点数
        r := dfs(root.Right)
        // 最大左右子树和结果比较和结果比较
        sum := l + r
        if sum > res {
            res = sum
        }

        // 返回左右子树中最大值加一
        if l > r {
            return l + 1
        } else {
            return r + 1
        }
    }

    // 调用递归函数
    dfs(root)

    // 返回结果
    return res
}
```
# 层次遍历
## [515. 在每个树行中找最大值](https://leetcode.cn/problems/find-largest-value-in-each-tree-row/description/)
```go
func largestValues(root *TreeNode) []int {
    // 节点为空直接返回
    if root == nil {
        return nil
    }


    // 初始化返回值和切片
    res, q := make([]int, 0), make([]*TreeNode, 0)
    q = append(q, root)


    // 切片不为空
    for len(q) > 0 {
        // 获取切片长度
        l := len(q)
        // 初始化当前层级最小值
        max := math.MinInt64
        // 遍历当前层级
        for i:=0; i<l; i++ {
            // 获取第一个节点
            t := q[0]
            // 从切片除去当前节点
            q = q[1:]


            // 当前节点和最大值比较看是否赋值
            v := t.Val
            if v > max {
                max = v
            }


            // 如果节点左子树存在，加入到下层切片中
            if t.Left != nil {
                q = append(q, t.Left)
            }


            // 如果节点右子树存在，加入到下层切片中
            if t.Right != nil {
                q = append(q, t.Right)
            }
        }


        // 当前层级最大值加入到结果切片中
        res = append(res, max)
    }


    // 返回结果
    return res
}
```
