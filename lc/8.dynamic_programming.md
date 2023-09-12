# [509. 斐波那契数](https://leetcode.cn/problems/fibonacci-number/description/)
```go
func fib(n int) int {
    // 递归终止条件
    if n == 0 {
        return 0
    }
    
    if n == 1 || n == 2 {
        return 1
    }

    // 计算当前结果并返回结果
    return fib(n-1) + fib(n-2)
}
```
```go
func fib(n int) int {
    // 初始化备忘录
    memo := make([]int, n+1)
    // 定义递归函数
    var dfs func(memo []int, n int) int
    dfs = func(memo []int, n int) int {
        // 终止条件
        if n == 0 {
            return 0
        }

        if n == 1 || n == 2 {
            return 1
        }

        // 备忘录存在，不用计算直接返回
        if memo[n] != 0 {
            return memo[n]
        }

        // 计算当前结果
        return dfs(memo, n - 1) + dfs(memo, n - 2)
    }

    // 返回结果
    return dfs(memo, n)
}
```
