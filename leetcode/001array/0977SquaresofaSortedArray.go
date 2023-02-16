package leetcode

// 给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
// 示例 1： 输入：nums = [-4,-1,0,3,10] 输出：[0,1,9,16,100] 解释：平方后，数组变为 [16,1,0,9,100]，排序后，数组变为 [0,1,9,16,100]
// 示例 2： 输入：nums = [-7,-3,2,3,11] 输出：[4,9,9,49,121]

// 采用双指针，从左右挑出最大值放到尾部，但是平方运算有点多做了还可以优化
// func sortedSquares(nums []int) []int {
// 	length := len(nums)
// 	res := make([]int, length)
// 	left, right := 0, length-1
// 	index := right
// 	for left <= right {
// 		leftSquare := nums[left] * nums[left]
// 		rightSquare := nums[right] * nums[right]
// 		if leftSquare >= rightSquare {
// 			res[index] = leftSquare
// 			left++
// 		} else {
// 			res[index] = rightSquare
// 			right--
// 		}
// 		index--
// 	}
// 	return res
// }
// 优化了平方，改为负号操作，应该是比平方要快的,但是这样的代码就更难理解了。
// 时间复杂度是O(n) 空间复杂度O(n)
func sortedSquares(nums []int) []int {
	length := len(nums)
	res := make([]int, length)
	left, right := 0, length-1
	index := right
	for left <= right {
		if -nums[left] >= nums[right] {
			res[index] = nums[left] * nums[left]
			left++
		} else {
			res[index] = nums[right] * nums[right]
			right--
		}
		index--
	}
	return res
}
