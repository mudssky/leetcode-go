package leetcode

// func sumOfLeftLeaves(root *TreeNode) int {
// 	var res int

// 	if root == nil {
// 		return 0
// 	}
// 	// 必须是父节点，才有左叶子节点
// 	if root.Left == nil && root.Right == nil {
// 		return 0
// 	}
// 	// 找到左叶子节点，需要通过它的父节点才能确定
// 	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
// 		res += root.Left.Val
// 	}
// 	res += sumOfLeftLeaves(root.Left)
// 	res += sumOfLeftLeaves(root.Right)
// 	return res
// }

// 迭代法，前序遍历
func sumOfLeftLeaves(root *TreeNode) int {
	stack := make([]*TreeNode, 0)
	if root == nil {
		return 0
	}
	stack = append(stack, root)
	result := 0

	for len(stack) != 0 {
		top := len(stack) - 1
		node := stack[top]
		stack = stack[:top]
		if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
			result += node.Left.Val
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return result
}
