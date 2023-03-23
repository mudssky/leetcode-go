package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }
// func minDepth(root *TreeNode) int {

// 	if root == nil {
// 		return 0
// 	}
// 	// 处理一边子树为空的特殊情况
// 	// 左子树为空，右子树不为空，返回右子树最小深度
// 	if root.Left == nil && root.Right != nil {
// 		return 1 + minDepth(root.Right)
// 	}
// 	// 返回左子树最小深度
// 	if root.Right == nil && root.Left != nil {
// 		return 1 + minDepth(root.Left)
// 	}
// 	return min(minDepth(root.Left), minDepth(root.Right)) + 1
// }

// 迭代法
func minDepth(root *TreeNode) int {
	dep := 0
	queue := make([]*TreeNode, 0)
	if root != nil {
		queue = append(queue, root)
	}
	for l := len(queue); l > 0; {
		dep++
		for ; l > 0; l-- {
			node := queue[0]
			if node.Left == nil && node.Right == nil {
				return dep
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			queue = queue[1:]
		}
		l = len(queue)
	}
	return dep
}
