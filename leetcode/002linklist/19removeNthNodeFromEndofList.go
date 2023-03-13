package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func removeNthFromEnd(head *ListNode, n int) *ListNode {
// 	dummy := &ListNode{}
// 	dummy.Next = head
// 	fast := dummy
// 	slow := dummy
// 	// 先让fast先走到n-1步
// 	for i := 0; i < n+1; i++ {
// 		// 这里不考虑n大于链表长度的情况
// 		fast = fast.Next
// 	}
// 	// 让fast走到末尾，此时slow在倒数n+1个元素
// 	for fast != nil {
// 		fast = fast.Next
// 		slow = slow.Next
// 	}
// 	slow.Next = slow.Next.Next
// 	return dummy.Next
// }

// 双指针也可以单次遍历完成
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	fast := head
	slow := dummy
	i := 1
	for fast != nil {
		fast = fast.Next
		if i > n {
			slow = slow.Next
		}
		i++
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}
