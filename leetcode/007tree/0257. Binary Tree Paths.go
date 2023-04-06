package leetcode

import "strconv"

// func binaryTreePaths(root *TreeNode) []string {
// 	res := make([]string, 0)
// 	var travel func(node *TreeNode, s string)
// 	travel = func(node *TreeNode, s string) {
// 		// 左右节点都为空，说明到达叶子节点了，
// 		// 路径找到，要结束递归了。
// 		if node.Left == nil && node.Right == nil {
// 			v := s + strconv.Itoa(node.Val)
// 			res = append(res, v)
// 			return
// 		}
// 		s = s + strconv.Itoa(node.Val) + "->"
// 		if node.Left != nil {
// 			travel(node.Left, s)
// 		}
// 		if node.Right != nil {
// 			travel(node.Right, s)
// 		}
// 	}
// 	travel(root, "")
// 	return res
// }

// 迭代法
func binaryTreePaths(root *TreeNode) []string {
	stack := []*TreeNode{}
	paths := make([]string, 0)
	res := make([]string, 0)
	if root != nil {
		stack = append(stack, root)
		paths = append(paths, "")
	}
	for len(stack) > 0 {
		top := len(stack) - 1
		node := stack[top]
		path := paths[top]
		stack = stack[:top]
		paths = paths[:top]
		// 迭代到叶子结点,继续回溯遍历
		if node.Left == nil && node.Right == nil {
			res = append(res, path+strconv.Itoa(node.Val))
			continue
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
			paths = append(paths, path+strconv.Itoa(node.Val)+"->")
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
			paths = append(paths, path+strconv.Itoa(node.Val)+"->")
		}
	}
	return res
}
