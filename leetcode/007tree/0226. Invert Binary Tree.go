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
//  递归法
// func invertTree(root *TreeNode) *TreeNode {
// 	if root == nil {
// 		return nil
// 	}
// 	root.Left, root.Right = root.Right, root.Left
// 	invertTree(root.Left)
// 	invertTree(root.Right)
// 	return root
// }

//  迭代法
// 用队列
func invertTree(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	if root == nil {
		return root
	}
	queue := list.New()
	node := root
	queue.PushBack(node)
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			e := queue.Remove(queue.Front()).(*TreeNode)
			e.Left, e.Right = e.Right, e.Left //交换
			if e.Left != nil {
				queue.PushBack(e.Left)
			}
			if e.Right != nil {
				queue.PushBack(e.Right)
			}
		}
	}
	return root

}
