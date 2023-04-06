package leetcode

import "container/list"

// func findBottomLeftValue(root *TreeNode) int {
// 	depth, res := 0, 0 // 初始化

// 	var dfs func(root *TreeNode, d int)
// 	dfs = func(root *TreeNode, d int) {
// 		if root == nil {
// 			return
// 		}
// 		// 因为先遍历左边，所以左边如果有值，右边的同层不会更新结果
// 		// 叶子节点是退出条件，并且需要大于已记录的最大深度才会更新
// 		if root.Left == nil && root.Right == nil && depth < d {
// 			depth = d
// 			res = root.Val
// 		}
// 		dfs(root.Left, d+1) // 隐藏回溯
// 		dfs(root.Right, d+1)
// 	}
// 	dfs(root, 1)
// 	return res
// }

// 迭代法，层序遍历
func findBottomLeftValue(root *TreeNode) int {
	var res int
	queue := list.New()

	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			// 每层记录第一个，最后一层遍历完，自然就是最后一层的第一个
			if i == 0 {
				res = node.Val
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}
	return res
}
