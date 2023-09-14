# [34. 在排序数组中查找元素的第一个和最后一个位置](https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/)
```go
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	l := 0
	r := len(nums) - 1
	for l < r {
		mid := (l + r) >> 1
		if target <= nums[mid] {
			r = mid
		} else {
			l = mid + 1
		}
	}

	if nums[l] != target {
		return []int{-1, -1}
	}
	t := l

	l = 0
	r = len(nums) - 1
	for l < r {
		mid := (l + r + 1) >> 1
		if target >= nums[mid] {
			l = mid
		} else {
			r = mid - 1
		}
	}

	if nums[r] != target {
		return []int{-1, -1}
	}

	return []int{t, r}
}

```
# [704. 二分查找](https://leetcode.cn/problems/binary-search/description/)
```go
func search(nums []int, target int) int {
    l := 0
    r := len(nums) - 1

    for l < r {
        m := (l + r) >> 1
        if target <= nums[m]{
            r = m
        } else {
            l = m + 1
        }
    }

    if nums[l] != target {
        return -1
    }

    return l
}
```
