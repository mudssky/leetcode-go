package leetcode

// 你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，
// 如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。

// 给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。
// func rob(nums []int) int {

// 	// dp[i]：考虑下标i（包括i）以内的房屋，最多可以偷窃的金额为dp[i]。
// 	dp := make([]int, len(nums))
// 	dp[0] = nums[0]
// 	if len(nums) < 2 {
// 		return nums[0]
// 	}
// 	dp[1] = max(nums[0], nums[1])
// 	for i := 2; i < len(nums); i++ {
// 		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
// 	}
// 	return dp[len(nums)-1]
// }

func rob(nums []int) int {
	length := len(nums)
	// dp[i]：偷到第i家房子，最多可以偷窃的金额为dp[i]。
	dp := make([]int, length+1)
	// dp[0]=0
	dp[1] = nums[0]
	for i := 2; i <= length; i++ {
		dp[i] = max(dp[i-2]+nums[i-1], dp[i-1])
	}
	return dp[length]
}
