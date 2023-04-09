package leetcode

func lowestCommonAncestorBST(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	for {
		if root.Val > p.Val && root.Val > q.Val {
			root = root.Left
		}
		if root.Val < p.Val && root.Val < q.Val {
			root = root.Right
		}
		// 节点的值在p,q之间
		if (root.Val-p.Val)*(root.Val-q.Val) <= 0 {
			return root
		}
	}
	// return nil
}
