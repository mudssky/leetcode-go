package leetcode

// 递归版本
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
		return root
	}
	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
		return root
	}
	// 找到节点的情况
	// 右子节点为空，此时左子节点补位，返回左子节点
	if root.Right == nil {
		return root.Left
	}
	if root.Left == nil {
		return root.Right
	}
	// 左右孩子节点都不为空，则将删除节点的左子树头结点（左孩子）放到删除节点的右子树的最左面节点的左孩子上，返回删除节点右孩子为新的根节点。
	// 这里root就是要删除的节点了。
	minnode := root.Right
	// 找到右节点里的最小值节点（最左）
	for minnode.Left != nil {
		minnode = minnode.Left
	}

	// 用了一种比较绕的交换法删除节点
	root.Val = minnode.Val
	root.Right = deleteNode1(root.Right)
	return root
}

// 交换法删除节点
func deleteNode1(root *TreeNode) *TreeNode {
	// 左子节点为空时，删除右子节点
	if root.Left == nil {
		pRight := root.Right
		root.Right = nil
		return pRight
	}
	root.Left = deleteNode1(root.Left)
	return root
}

// 迭代版本
// func deleteOneNode(target *TreeNode) *TreeNode {
// 	if target == nil {
// 		return target
// 	}
// 	if target.Right == nil {
// 		return target.Left
// 	}
// 	cur := target.Right
// 	for cur.Left != nil {
// 		cur = cur.Left
// 	}
// 	cur.Left = target.Left
// 	return target.Right
// }
// func deleteNode(root *TreeNode, key int) *TreeNode {
// 	// 特殊情况处理
// 	if root == nil {
// 		return root
// 	}
// 	cur := root
// 	var pre *TreeNode
// 	for cur != nil {
// 		if cur.Val == key {
// 			break
// 		}
// 		pre = cur
// 		if cur.Val > key {
// 			cur = cur.Left
// 		} else {
// 			cur = cur.Right
// 		}
// 	}
// 	if pre == nil {
// 		return deleteOneNode(cur)
// 	}
// 	// pre 要知道是删除左孩子还有右孩子
// 	if pre.Left != nil && pre.Left.Val == key {
// 		pre.Left = deleteOneNode(cur)
// 	}
// 	if pre.Right != nil && pre.Right.Val == key {
// 		pre.Right = deleteOneNode(cur)
// 	}
// 	return root
// }

func deleteNodeSwap(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	// 找到目标节点
	if root.Val == key {
		// 如果在交换之前碰到这个节点(只会是叶子节点)，起到删除的作用
		// 交换之后，则是在右子树的最左节点碰到，同样起到删除节点的作用。
		if root.Right == nil {
			return root.Left
		}
		// 找到最左子节点
		cur := root.Right
		for cur.Left != nil {
			cur = cur.Left
		}
		// 交换目标节点和最左子节点
		root.Val, cur.Val = cur.Val, root.Val
	}
	root.Left = deleteNodeSwap(root.Left, key)
	root.Right = deleteNodeSwap(root.Right, key)
	return root
}
