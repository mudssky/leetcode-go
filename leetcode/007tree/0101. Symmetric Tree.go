package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//  递归法，比较容易理解
// 但是调用次数多会有栈溢出的问题
// func defs(left *TreeNode, right *TreeNode) bool {
// 	if left == nil && right == nil {
// 		return true
// 	}
// 	if left == nil || right == nil {
// 		return false
// 	}
// 	if left.Val != right.Val {
// 		return false
// 	}
// 	return defs(left.Left, right.Right) && defs(right.Left, left.Right)
// }
// func isSymmetric(root *TreeNode) bool {

// 	return defs(root.Left, root.Right)
// }

// 迭代
func isSymmetric(root *TreeNode) bool {
	var queue []*TreeNode
	if root != nil {
		queue = append(queue, root.Left, root.Right)
	}
	for len(queue) > 0 {
		left := queue[0]
		right := queue[1]
		queue = queue[2:]
		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}
		queue = append(queue, left.Left, right.Right, right.Left, left.Right)
	}
	return true
}
