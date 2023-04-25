package leetcode

// 给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。

// 子序列是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。

func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	// dp[i]表示i之前包括i的以nums[i]结尾的最长递增子序列的长度
	// 这样方便我们进行比较nums[i]和nums[j]
	dp := make([]int, n)
	// 初始所有最长子序列为1，因为只要字符串有长度，至少会有1
	for i := range dp {
		dp[i] = 1
	}
	res := 0
	// 因为dp[i]不是记录普遍的最长长度，而是以nums[i]结尾的最长子序列，所以我们要在途中记录最大值
	for i := 1; i < n; i++ {
		// 遍历子序列
		for j := 0; j < i; j++ {
			// nums[i]更大的情况，最长子序列才有增长的空间
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
			if dp[i] > res {
				res = dp[i]
			}
		}
	}
	return res
}
