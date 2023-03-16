package leetcode

// func postorderTraversal(root *TreeNode) []int {
// 	res := []int{}
// 	if root == nil {
// 		return res
// 	}

// 	res = append(res, postorderTraversal(root.Left)...)
// 	res = append(res, postorderTraversal(root.Right)...)
// 	res = append(res, root.Val)
// 	return res
// }

// 迭代法
// 因为后序遍历 左右中
// 就是先序遍历 中左右,调换 中右左,然后反转就可以
func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	stack := make([]*TreeNode, 1)
	stack[0] = root
	for len(stack) > 0 {
		lastIndex := len(stack) - 1
		node := stack[lastIndex]
		res = append(res, node.Val)
		// 出栈
		stack = stack[:lastIndex]
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}

	reverse(res)
	return res
}

func reverse(a []int) {
	l, r := 0, len(a)-1
	for l < r {
		a[l], a[r] = a[r], a[l]
		l, r = l+1, r-1
	}
}
