package leetcode

// func insertIntoBST(root *TreeNode, val int) *TreeNode {

// 	if root == nil {
// 		root = &TreeNode{Val: val}
// 		return root
// 	}
// 	if root.Val > val {
// 		root.Left = insertIntoBST(root.Left, val)
// 	} else {
// 		root.Right = insertIntoBST(root.Right, val)
// 	}
// 	return root
// }

// 迭代法
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	node := root
	var pnode *TreeNode
	for node != nil {
		if val > node.Val {
			pnode = node
			node = node.Right
		} else {
			pnode = node
			node = node.Left
		}
	}
	if val > pnode.Val {
		pnode.Right = &TreeNode{Val: val}
	} else {
		pnode.Left = &TreeNode{Val: val}
	}
	return root
}
