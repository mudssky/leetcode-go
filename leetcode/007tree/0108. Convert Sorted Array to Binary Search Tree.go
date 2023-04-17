package leetcode

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 { //终止条件，最后数组为空则可以返回
		return nil
	}
	i := len(nums) / 2
	root := &TreeNode{Val: nums[i]}

	root.Left = sortedArrayToBST(nums[:i])
	root.Right = sortedArrayToBST(nums[i+1:])

	return root
}
