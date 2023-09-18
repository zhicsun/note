# 模板
## 快速排序
```go
package main

func quickSort(s []int, l int, r int) {
	// 递归结束条件
	if l >= r {
		return
	}

	// 递归处理使左侧小于中间值, 右侧大于中间值
	left, right, x := l-1, r+1, s[(l+r)>>1]
	for left < right {
		// 找到第一个大于中间值的下标
		for {
			left++
			if s[left] >= x {
				break
			}
		}

		// 找到第一个小于中间值下标
		for {
			right--
			if s[right] <= x {
				break
			}
		}

		// 左右侧值交换
		if left < right {
			s[left], s[right] = s[right], s[left]
		}
	}
	
	// 递归处理左右两侧值
	quickSort(s, l, right)
	quickSort(s, right+1, r)
}

```