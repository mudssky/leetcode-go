package leetcode

import "math"

// 中序遍历下，输出的二叉搜索树节点的数值是有序序列
// 可以先中序遍历出数组，然后判断数组是否为升序
// func isValidBST(root *TreeNode) bool {
// 	list := inorderTraversal(root)
// 	for i := 0; i < len(list)-1; i++ {
// 		if list[i] >= list[i+1] {
// 			return false
// 		}
// 	}
// 	return true
// }

func isValidBST(root *TreeNode) bool {
	// 二叉搜索树也可以是空树
	if root == nil {
		return true
	}
	// 由题目中的数据限制可以得出min和max
	return check(root, math.MinInt64, math.MaxInt64)
}

func check(node *TreeNode, min, max int64) bool {
	if node == nil {
		return true
	}

	if min >= int64(node.Val) || max <= int64(node.Val) {
		return false
	}
	// 分别对左子树和右子树递归判断，如果左子树和右子树都符合则返回true
	return check(node.Left, min, int64(node.Val)) && check(node.Right, int64(node.Val), max)
}
