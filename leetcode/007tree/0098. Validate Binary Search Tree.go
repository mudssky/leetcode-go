package leetcode

// 中序遍历下，输出的二叉搜索树节点的数值是有序序列
// 可以先中序遍历出数组，然后判断数组是否为升序
func isValidBST(root *TreeNode) bool {
	list := inorderTraversal(root)
	for i := 0; i < len(list)-1; i++ {
		if list[i] >= list[i+1] {
			return false
		}
	}
	return true
}
