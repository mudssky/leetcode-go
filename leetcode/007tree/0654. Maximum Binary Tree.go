package leetcode

// 给定一个不含重复元素的整数数组。一个以此数组构建的最大二叉树定义如下：

// 二叉树的根是数组中的最大元素。
// 左子树是通过数组中最大值左边部分构造出的最大二叉树。
// 右子树是通过数组中最大值右边部分构造出的最大二叉树。
// 通过给定的数组构建最大二叉树，并且输出这个树的根节点。

func constructMaximumBinaryTree(nums []int) *TreeNode {

	if len(nums) == 0 {
		return nil
	}
	// 找到最大值
	index := findMax(nums)
	// 构造二叉树
	root := &TreeNode{
		Val:   nums[index],
		Left:  constructMaximumBinaryTree(nums[:index]),   //左半边
		Right: constructMaximumBinaryTree(nums[index+1:]), //右半边
	}
	return root
}
func findMax(nums []int) (index int) {
	for i, v := range nums {
		if nums[index] < v {
			index = i
		}
	}
	return
}
