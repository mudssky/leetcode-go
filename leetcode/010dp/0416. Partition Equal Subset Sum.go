package leetcode

// 给定一个只包含正整数的非空数组。是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
// 注意: 每个数组中的元素不会超过 100 数组的大小不会超过 200

// 分割等和子集 动态规划
// 时间复杂度O(n^2) 空间复杂度O(n)
func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	// 如果 nums 的总和为奇数则不可能平分成两个子集
	if sum%2 == 1 {
		return false
	}

	target := sum / 2
	dp := make([]int, target+1)

	for _, num := range nums {
		for j := target; j >= num; j-- {
			if dp[j] < dp[j-num]+num {
				dp[j] = dp[j-num] + num
			}
		}
	}
	return dp[target] == target
}
