package leetcode

// 给定一个有相同值的二叉搜索树（BST），找出 BST 中的所有众数（出现频率最高的元素）。

func findMode(root *TreeNode) []int {
	res := make([]int, 0)
	count := 1
	max := 1
	var prev *TreeNode
	var travel func(node *TreeNode)
	travel = func(node *TreeNode) {
		if node == nil {
			return
		}
		travel(node.Left)
		if prev != nil && prev.Val == node.Val {
			count++
		} else {
			count = 1
		}
		if count >= max {
			if count > max && len(res) > 0 {
				res = []int{node.Val}
			} else {
				res = append(res, node.Val)
			}
			max = count
		}
		prev = node
		travel(node.Right)
	}
	travel(root)
	return res
}
