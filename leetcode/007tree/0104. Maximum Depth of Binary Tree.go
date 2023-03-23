package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//  递归法
// func maxDepth(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}
// 	leftMax := maxDepth(root.Left) + 1
// 	rightMax := maxDepth(root.Right) + 1

// 	if leftMax > rightMax {
// 		return leftMax
// 	}
// 	return rightMax
// }

// 迭代法
func maxDepth(root *TreeNode) int {
	level := 0
	queue := make([]*TreeNode, 0)
	if root != nil {
		queue = append(queue, root)
	}
	for l := len(queue); l > 0; {
		for ; l > 0; l-- {
			node := queue[0]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			queue = queue[1:]
		}
		level++
		l = len(queue)
	}
	return level
}
