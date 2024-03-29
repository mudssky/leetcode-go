package leetcode

import "github.com/mudssky/leetcode-go/structures"

type TreeNode = structures.TreeNode

func rob3(root *TreeNode) int {
	res := robTree(root)
	return max(res[0], res[1])
}

func robTree(cur *TreeNode) []int {
	if cur == nil {
		return []int{0, 0}
	}
	// 后序遍历
	left := robTree(cur.Left)
	right := robTree(cur.Right)

	// 考虑去偷当前的屋子,则子节点不偷
	robCur := cur.Val + left[0] + right[0]
	// 考虑不去偷当前的屋子，则计算子节点所有情况的最大值
	notRobCur := max(left[0], left[1]) + max(right[0], right[1])

	// 注意顺序：0:不偷，1:去偷
	return []int{notRobCur, robCur}
}
