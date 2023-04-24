package leetcode

// 你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，
// 这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。

// 给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，能够偷窃到的最高金额。

// 成环以后，可以分为三种情况
// 1.只考虑头部
// 2.不考虑头尾，取中间
// 3.只考虑尾部
// 其中1，3的范围包括了2的范围，所以我们只要计算2，3两种情况
func rob2(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}
	result1 := rob(nums[0 : len(nums)-1])
	result2 := rob(nums[1:])
	return max(result1, result2)
}

// 偷盗指定的范围,范围用的是nums的下标
// func robRange(nums []int) int {
// 	length := len(nums)
// 	// dp[i]：偷到第i家房子，最多可以偷窃的金额为dp[i]。
// 	dp := make([]int, length+1)
// 	// dp[0]=0
// 	dp[1] = nums[0]
// 	for i := 2; i <= length; i++ {
// 		dp[i] = max(dp[i-2]+nums[i-1], dp[i-1])
// 	}
// 	return dp[length]
// }
