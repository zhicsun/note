# [21. 合并两个有序链表](https://leetcode.cn/problems/merge-two-sorted-lists/description/)
```go
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    // 终止条件和返回结果
    if list1 == nil {
        return list2
    }

    if list2 == nil {
        return list1
    }

    // 从小到大排序
    if list1.Val <= list2.Val {
        list1.Next = mergeTwoLists(list1.Next, list2)
        return list1
    } else {
        list2.Next = mergeTwoLists(list1, list2.Next)
        return list2
    }
}
```
# [206. 反转链表](https://leetcode.cn/problems/reverse-linked-list/)
```go
func reverseList(head *ListNode) *ListNode {
    // 递归终止条件，返回最后一个节点
    if head == nil || head.Next == nil {
        return head
    }

    t := reverseList(head.Next)

    // 离开节点后处理
    head.Next.Next = head
    head.Next = nil

    return t
}
```
