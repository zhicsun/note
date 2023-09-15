# 双指针

## 数组

### [5. 最长回文子串](https://leetcode.cn/problems/longest-palindromic-substring/)

```go
package main

func longestPalindrome(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		// 以 s[i] 为中心的最长回文子串
		s1 := palindrome(s, i, i)
		// 以 s[i] 和 s[i+1] 为中心的最长回文子串
		s2 := palindrome(s, i, i+1)
		if len(res) > len(s1) {
			res = res
		} else {
			res = s1
		}
		if len(res) > len(s2) {
			res = res
		} else {
			res = s2
		}
	}
	return res
}

func palindrome(s string, l int, r int) string {
	// 防止索引越界
	for l >= 0 && r < len(s) && s[l] == s[r] {
		// 向两边展开f
		l--
		r++
	}
	// 返回以 s[l] 和 s[r] 为中心的最长回文串
	return s[l+1 : r]
}

```

### [26. 删除有序数组中的重复项](https://leetcode.cn/problems/remove-duplicates-from-sorted-array/)

```go
package main

func removeDuplicates(nums []int) int {
	// 数量小于等于两个直接返回
	if len(nums) <= 1 {
		return len(nums)
	}

	// 快指针不断往后移动，快指针不等于慢指针，慢指针往后移动一位，快指针的值赋给慢指针
	f, s := 1, 0
	for ; f < len(nums); f++ {
		if nums[f] != nums[s] {
			s++
			nums[s] = nums[f]
		}
	}

	// 返回结果
	return s + 1
}

```

### [27. 移除元素](https://leetcode.cn/problems/remove-element/)

```go
package main

func removeElement(nums []int, val int) int {
	// 初始化快慢指针
	f, s := 0, 0
	for ; f < len(nums); f++ {
		// 快指针的值不等于删除值，慢指等于当前值并移动慢指针
		if nums[f] != val {
			nums[s] = nums[f]
			s++
		}
	}

	// 返回结果
	return s
}

```

### [167. 两数之和 II - 输入有序数组](https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/)

```go
package main

func twoSum(numbers []int, target int) []int {
	// 初始化左右两个指针
	l, r := 0, len(numbers)-1

	// 左指针小于右指针
	for l < r {
		// 左右指针值相加等于目标值，返回结果
		sum := numbers[l] + numbers[r]
		if sum == target {
			return []int{l + 1, r + 1}
		}

		// 左右指针值大于目标值，右指针移动
		if sum > target {
			r--
		}

		// 左右指针值小于目标值，左指针移动
		if sum < target {
			l++
		}
	}

	// 返回结果
	return []int{-1, -1}
}

```

### [283. 移动零](https://leetcode.cn/problems/move-zeroes/description/)

```go
package main

func moveZeroes(nums []int) {
	// 初始化快慢指针
	f, s := 0, 0
	for ; f < len(nums); f++ {
		// 如果快指针不等于，交互快慢指针，慢指针往后移动
		if nums[f] != 0 {
			nums[s], nums[f] = nums[f], nums[s]
			s++
		}
	}
}

```

### [344. 反转字符串](https://leetcode.cn/problems/reverse-string/description/)

```go
package main

func reverseString(s []byte) {
	// 初始化左右指针
	l, r := 0, len(s)-1

	// 左指针小于右指针
	for l < r {
		// 交互左右指针的值
		s[l], s[r] = s[r], s[l]

		// 左指针相加
		l++
		// 右指针相减
		r--
	}
}

```

## 链表

### [19. 删除链表的倒数第 N 个结点](https://leetcode.cn/problems/remove-nth-node-from-end-of-list/description/)

```go
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 构建虚拟节点
	dummy := &ListNode{Next: head}

	// 要想删除倒数第 n 节点，要知道 n + 1 节点，先移动 n + 1 次 
	f := dummy
	for i := 0; i <= n; i++ {
		f = f.Next
	}

	// 快指针移动到终点，慢指针指向倒数 n + 1 位
	s := dummy
	for f != nil {
		f = f.Next
		s = s.Next
	}

	// 删除倒数第 n 位
	s.Next = s.Next.Next

	// 返回结果
	return dummy.Next
}

```

### [21. 合并两个有序链表](https://leetcode.cn/problems/merge-two-sorted-lists/description/)

```go
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 构造虚拟节点
	dummy := &ListNode{}
	// 复制虚拟节点
	current := dummy
	// 遍历同等长度节点，递增构造
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	// 拼接 list1 链表 
	for list1 != nil {
		current.Next = &ListNode{Val: list1.Val}
		list1 = list1.Next
		current = current.Next
	}

	// 拼接 list2 链表
	for list2 != nil {
		current.Next = &ListNode{Val: list2.Val}
		list2 = list2.Next
		current = current.Next
	}

	// 返回结果
	return dummy.Next
}

```

### [23. 合并 K 个升序链表](https://leetcode.cn/problems/merge-k-sorted-lists/description/)

```go
package main

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	// 初始化优先队列
	q := &priorityQueue{}
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			heap.Push(q, lists[i])
		}
	}

	// 初始化虚拟节点和构造节点
	dummy := new(ListNode)
	p := dummy
	// 循环有先队列，每次取最小值，填入构造节点中
	for q.Len() > 0 {
		top := heap.Pop(q).(*ListNode)
		if top.Next != nil {
			heap.Push(q, top.Next)
		}
		p.Next, p = top, top
	}

	// 返回结果
	return dummy.Next
}

// 实现优先队列
type priorityQueue []*ListNode

func (h priorityQueue) Len() int {
	return len(h)
}

func (h priorityQueue) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h priorityQueue) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *priorityQueue) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *priorityQueue) Pop() interface{} {
	l := len(*h) - 1
	x := (*h)[l]
	*h = (*h)[0:l]
	return x
}

```

### [83. 删除排序链表中的重复元素](https://leetcode.cn/problems/remove-duplicates-from-sorted-list/)

```go
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	// 为空直接返回
	if head == nil {
		return nil
	}

	// 初始化快慢指针
	f, s := head, head
	for f != nil {
		// 快慢指针不同，慢指针接快指针当前点，慢指针往后移动
		if s.Val != f.Val {
			s.Next = f
			s = s.Next
		}
		// 移动快指针
		f = f.Next
	}

	// 除去多余的值
	s.Next = nil

	// 返回结果
	return head
}

```

### [86. 分隔链表](https://leetcode.cn/problems/partition-list/description/)

```go
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	// 构造虚拟节点
	dummy1 := &ListNode{}
	dummy2 := &ListNode{}

	// 复制虚拟节点
	p1 := dummy1
	p2 := dummy2

	// 构造小于和大于等于指定值的两个链表
	for head != nil {
		if head.Val < x {
			p1.Next = head
			p1 = p1.Next
		} else {
			p2.Next = head
			p2 = p2.Next
		}
		// 断开当前节点连接
		temp := head.Next
		head.Next = nil
		head = temp
	}

	// 两个链表连接
	p1.Next = dummy2.Next

	//返回结果
	return dummy1.Next
}

```

### [141. 环形链表](https://leetcode.cn/problems/linked-list-cycle/)

```go
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	f := head
	s := head

	for f != nil && f.Next != nil {
		f = f.Next.Next
		s = s.Next
		if f == s {
			return true
		}
	}

	return false
}

```

### [142. 环形链表 II](https://leetcode.cn/problems/linked-list-cycle-ii/)

slow 在相遇点走过的距离: a+b

fast 在相遇点走过的距离: a+n(b+c)+b=a+(n+1)b+nc

fast 在相遇点走过的距离是 slow 的 2 倍: a+(n+1)b+nc=2(a+b) ==> a=c+(n−1)(b+c)

```go
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	// 构建快慢指针
	f := head
	s := head

	// 获取到相遇点
	for f != nil && f.Next != nil {
		f = f.Next.Next
		s = s.Next
		if f == s {
			break
		}
	}

	// 相遇点往后没数据
	if f == nil || f.Next == nil {
		return nil
	}

	// 获取环入口
	s = head
	for s != f {
		f = f.Next
		s = s.Next
	}

	// 返回结果
	return s
}

```

### [160. 相交链表](https://leetcode.cn/problems/intersection-of-two-linked-lists/)

```go
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 构造虚拟节点
	p1 := headA
	p2 := headB

	// 寻找相同节点，到尾部相互交换起点，相交点为返回结果
	for p1 != p2 {
		if p1 == nil {
			p1 = headB
		} else {
			p1 = p1.Next
		}

		if p2 == nil {
			p2 = headA
		} else {
			p2 = p2.Next
		}
	}

	// 返回结果
	return p2
}

```

### [876. 链表的中间结点](https://leetcode.cn/problems/middle-of-the-linked-list/)

```go
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	f := head
	s := head

	for f != nil && f.Next != nil {
		f = f.Next.Next
		s = s.Next
	}

	return s
}

```
