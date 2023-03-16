package leetcode

import (
	"github.com/mudssky/leetcode-go/structures"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode = structures.TreeNode

// 递归法
// func preorderTraversal(root *TreeNode) []int {
// 	var res []int
// 	if root == nil {
// 		return []int{}
// 	}
// 	res = append(res, root.Val)
// 	res = append(res, preorderTraversal(root.Left)...)
// 	res = append(res, preorderTraversal(root.Right)...)
// 	return res
// }

// 迭代法
// 要用到栈
// func preorderTraversal(root *TreeNode) []int {
// 	res := []int{}
// 	if root == nil {
// 		return res
// 	}
// 	stack := make([]*TreeNode, 1)
// 	stack[0] = root
// 	for len(stack) > 0 {
// 		lastIndex := len(stack) - 1
// 		node := stack[lastIndex]
// 		res = append(res, node.Val)
// 		// 出栈
// 		stack = stack[:lastIndex]
// 		if node.Right != nil {
// 			stack = append(stack, node.Right)
// 		}
// 		if node.Left != nil {
// 			stack = append(stack, node.Left)
// 		}
// 	}
// 	return res
// }

// 标记法
// 之前迭代法的前中后序遍历很不一致，因为中序遍历的时候，访问节点（遍历节点顺序）和处理节点（元素放入结果集）不一致。
// 所以我们可以在处理节点入栈后放入一个空指针作为标记
func preorderTraversal(root *TreeNode) []int {

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
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		// 先序遍历，把中节点放在栈顶
		stack = append(stack, node)
		stack = append(stack, nil)
	}

	return res
}
