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

## 归并排序

```go
package main

func mergeSort(s []int, l int, r int) {
	// 递归结束条件
	if l >= r {
		return
	}

	// 获取中间节点
	mid := (l + r) >> 1

	// 递归处理中间两侧数据
	mergeSort(s, l, mid)
	mergeSort(s, mid+1, r)

	// 使中间节点两侧数据有序
	var t []int
	left, right := l, mid+1
	for left <= mid && right <= r {
		if s[left] <= s[right] {
			t = append(t, s[left])
			left++
		} else {
			t = append(t, s[right])
			right++
		}
	}

	// 处理左侧剩余数据
	for left <= mid {
		t = append(t, s[left])
		left++
	}

	// 处理右侧剩余数据
	for right <= r {
		t = append(t, s[right])
		right++
	}

	// 更改原来数据使之有序
	for left, right = l, 0; left <= r; left, right = left+1, right+1 {
		s[left] = t[right]
	}
}

```