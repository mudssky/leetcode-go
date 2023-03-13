package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//  最简单的是用map两次遍历
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	nodeMap := make(map[*ListNode]bool, 100)
	for cur := headA; cur != nil; cur = cur.Next {
		nodeMap[cur] = true
	}
	for cur := headB; cur != nil; cur = cur.Next {
		if _, ok := nodeMap[cur]; ok {
			return cur
		}
	}
	return nil
}

// 双指针法，不用map，空间复杂度低

// func getIntersectionNode(headA, headB *ListNode) *ListNode {
//     curA := headA
//     curB := headB
//     lenA, lenB := 0, 0
//     // 求A，B的长度
//     for curA != nil {
//         curA = curA.Next
//         lenA++
//     }
//     for curB != nil {
//         curB = curB.Next
//         lenB++
//     }
//     var step int
//     var fast, slow *ListNode
//     // 请求长度差，并且让更长的链表先走相差的长度
//     if lenA > lenB {
//         step = lenA - lenB
//         fast, slow = headA, headB
//     } else {
//         step = lenB - lenA
//         fast, slow = headB, headA
//     }
//     for i:=0; i < step; i++ {
//         fast = fast.Next
//     }
//     // 遍历两个链表遇到相同则跳出遍历
//     for fast != slow {
//         fast = fast.Next
//         slow = slow.Next
//     }
//     return fast
// }
