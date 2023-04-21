package leetcode

// 给你一个整数数组 cost ，其中 cost[i] 是从楼梯第 i 个台阶向上爬需要支付的费用。
// 一旦你支付此费用，即可选择向上爬一个或者两个台阶。
// 你可以选择从下标为 0 或下标为 1 的台阶开始爬楼梯。

func minCostClimbingStairs(cost []int) int {
	// dp记录爬到第i个台阶的最小花费
	dp := make([]int, len(cost)+1)
	// 第一步默认不花体力
	dp[0] = 0
	dp[1] = 0
	for i := 2; i <= len(cost); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[len(cost)]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
