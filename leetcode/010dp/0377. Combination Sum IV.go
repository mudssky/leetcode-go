package leetcode

// 给定一个由正整数组成且不存在重复数字的数组，找出和为给定目标正整数的组合的个数。

func combinationSum4(nums []int, target int) int {
	// dp[i]: 凑成目标正整数为i的排列个数为dp[i]
	dp := make([]int, target+1)
	// 初始化
	dp[0] = 1
	// 先遍历背包(target)，在遍历物品(nums)，这样得到物品的排列
	for j := 0; j <= target; j++ {
		for i := 0; i < len(nums); i++ {
			if nums[i] <= j {
				// 可以放入的情况，就放入dp[j-nums[i]]表示nums[i]不放入时的组合数
				dp[j] += dp[j-nums[i]]
			}
		}
	}
	return dp[target]

}
