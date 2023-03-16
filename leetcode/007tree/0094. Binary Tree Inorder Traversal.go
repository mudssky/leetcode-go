package leetcode

// func inorderTraversal(root *TreeNode) []int {
// 	var res []int
// 	if root == nil {
// 		return []int{}
// 	}
// 	res = append(res, inorderTraversal(root.Left)...)
// 	res = append(res, root.Val)
// 	res = append(res, inorderTraversal(root.Right)...)
// 	return res
// }

// 中序遍历
// func inorderTraversal(root *TreeNode) []int {
// 	res := []int{}
// 	if root == nil {
// 		return res
// 	}
// 	res = append(res, inorderTraversal(root.Left)...)
// 	res = append(res, root.Val)
// 	res = append(res, inorderTraversal(root.Right)...)
// 	return res
// }

// func inorderTraversal(root *TreeNode) []int {
// 	res := []int{}
// 	stack := make([]*TreeNode, 0)
// 	cur := root

// 	for cur != nil || len(stack) > 0 {
// 		if cur != nil {
// 			stack = append(stack, cur)

// 			cur = cur.Left // 左
// 		} else {
// 			lastIndex := len(stack) - 1
// 			top := stack[lastIndex]
// 			stack = stack[:lastIndex]
// 			cur = top

// 			res = append(res, cur.Val) //中
// 			cur = cur.Right            //右
// 		}
// 	}
// 	return res
// }

// 标记法
func inorderTraversal(root *TreeNode) []int {

	res := []int{}
	if root == nil {
		return res
	}

	stack := make([]*TreeNode, 1)
	stack[0] = root
	for len(stack) > 0 {
		lastIndex := len(stack) - 1
		node := stack[lastIndex]
		stack = stack[:lastIndex]

		// 空的情况表示需要处理中间节点
		if node == nil {
			lastIndex := len(stack) - 1
			node := stack[lastIndex]
			stack = stack[:lastIndex]
			res = append(res, node.Val)
			continue
		}

		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		// 先序遍历，把中节点放在栈顶
		stack = append(stack, node)
		stack = append(stack, nil)
		if node.Left != nil {
			stack = append(stack, node.Left)
		}

	}

	return res
}
