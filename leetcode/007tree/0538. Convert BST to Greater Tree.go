package leetcode

// 采用反中序遍历，相当于升序的数组，从后往前累加。
func convertBST(root *TreeNode) *TreeNode {
	// 记录前一个节点
	pre := 0
	// 递归遍历并执行累加操作的函数
	var traversal func(cur *TreeNode)
	traversal = func(cur *TreeNode) {
		if cur == nil {
			return
		}
		// 右
		traversal(cur.Right)
		// 中
		cur.Val += pre
		pre = cur.Val
		// 左
		traversal(cur.Left)
	}
	traversal(root)
	return root
}
