package leetcode

// 给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
// 设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

// func maxProfit(prices []int) int {
// 	res := 0
// 	for i := 1; i < len(prices); i++ {
// 		if prices[i] > prices[i-1] {
// 			res += prices[i] - prices[i-1]
// 		}
// 	}
// 	return res
// }

// 也可以用动态规划法
func maxProfit(prices []int) int {
	// dp[i][0] 表示第i天持有股票所得现金。
	// dp[i][1] 表示第i天不持有股票所得最多现金
	dp := make([][]int, len(prices))
	// 这个只是分配连续空间而已
	status := make([]int, len(prices)*2)
	for i := range dp {
		dp[i] = status[:2]
		status = status[2:]
	}
	dp[0][0] = -prices[0]

	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}

	return dp[len(prices)-1][1]
}
