package leetcode

func buildTree(inorder []int, postorder []int) *TreeNode {
	hash := make(map[int]int)
	for i, v := range inorder { // 用map保存中序序列的数值对应位置
		hash[v] = i

	}
	var rebuild func(inorder []int, postorder []int, rootIdx int, l, r int) *TreeNode
	rebuild = func(inorder []int, postorder []int, rootIdx int, l, r int) *TreeNode {
		// l,r是分割的中序数组的左右下标。
		if l > r { // 说明没有元素，返回空树
			return nil
		}
		if l == r { // 只剩唯一一个元素，直接返回
			return &TreeNode{Val: inorder[l]}
		}
		rootV := postorder[rootIdx]   // 根据后序数组找到根节点的值
		rootIn := hash[rootV]         // 找到根节点在对应的中序数组中的位置
		root := &TreeNode{Val: rootV} // 构造根节点
		// 重建左节点和右节点
		root.Left = rebuild(inorder, postorder, rootIdx-(r-rootIn)-1, l, rootIn-1)
		root.Right = rebuild(inorder, postorder, rootIdx-1, rootIn+1, r)
		return root
	}
	// 以左闭右闭的原则进行切分
	root := rebuild(inorder, postorder, len(postorder)-1, 0, len(inorder)-1)
	return root
}
