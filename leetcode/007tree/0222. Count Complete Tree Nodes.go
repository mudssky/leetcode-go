package leetcode

import "container/list"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// func countNodes(root *TreeNode) int {
// 	var res int
// 	if root != nil {
// 		res = 1 + countNodes(root.Left) + countNodes(root.Right)

// 	}
// 	return res
// }
// 利用完全二叉树特性的解法
// 如果子树是满二叉树的情况，我们可以用2^n-1的公式计算节点数量，可以节省遍历的时间
// 完全二叉树，是满二叉树的情况，那么向右计算和向左计算的深度是相等的
// func countNodes(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}
// 	// 计算判断是否满的完全二叉树
// 	leftH, rightH := 0, 0
// 	leftNode := root.Left
// 	rightNode := root.Right

// 	for leftNode != nil {
// 		leftNode = leftNode.Left
// 		leftH++
// 	}
// 	for rightNode != nil {
// 		rightNode = rightNode.Right
// 		rightH++
// 	}
// 	// 如果是满的完全二叉树，套用公式计算
// 	if leftH == rightH {
// 		return (2 << leftH) - 1
// 	}
// 	// 不是满的完全二叉树，递归计算
// 	return countNodes(root.Left) + countNodes(root.Right) + 1
// }

// 下面试试迭代法层序遍历
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := list.New()
	q.PushBack(root)
	res := 0
	for q.Len() > 0 {
		n := q.Len()
		for i := 0; i < n; i++ {
			node := q.Remove(q.Front()).(*TreeNode)
			if node.Left != nil {
				q.PushBack(node.Left)
			}
			if node.Right != nil {
				q.PushBack(node.Right)
			}
			res++
		}
	}
	return res
}
