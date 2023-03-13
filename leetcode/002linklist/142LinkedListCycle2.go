package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//  用map就比较简单
// func detectCycle(head *ListNode) *ListNode {
// 	nodeMap := make(map[*ListNode]bool, 100)
// 	for cur := head; cur != nil; cur = cur.Next {
// 		if _, ok := nodeMap[cur]; ok {
// 			return cur
// 		}
// 		nodeMap[cur] = true
// 	}
// 	return nil
// }

// 使用O(1)空间解决
// 需要一些数学知识，首先只要设置两个指针，一快一慢，如果慢的最后和快的相遇，说明有环，有环的情况下，快指针和慢指针的相对速度是1，相当于快指针追上慢指针一个环的距离。
// 然后经过数学证明，可以知道，有环的情况，设置两个指针，一个从相遇位置出发，一个从链表起点出发，相遇的位置就是环的起点。
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			for slow != head {
				slow = slow.Next
				head = head.Next
			}
			return head
		}
	}
	return nil
}
