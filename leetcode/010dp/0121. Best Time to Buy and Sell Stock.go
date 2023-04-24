package leetcode

// 贪心解法
// func maxProfit(prices []int) int {
// 	low := prices[0]
// 	res := math.MinInt
// 	for i := 0; i < len(prices); i++ {
// 		// 计算最小价格
// 		low = min(low, prices[i])
// 		// 计算最大利润
// 		res = max(res, prices[i]-low)
// 	}
// 	return res
// }

// 动态规划
func maxProfit(prices []int) int {
	length := len(prices)
	if length == 0 {
		return 0
	}
	// dp[i][0] 表示第i天持有股票所得最多现金，指的是收益，因此不卖出的情况不会改变
	// dp[i][1] 表示第i天不持有股票所得最多现金
	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	for i := 1; i < length; i++ {
		dp[i][0] = max(dp[i-1][0], -prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}
	return dp[length-1][1]
}

// 如果第i天持有股票即dp[i][0]， 那么可以由两个状态推出来

// 第i-1天就持有股票，那么就保持现状，所得现金就是昨天持有股票的所得现金 即：dp[i - 1][0]
// 第i天买入股票，所得现金就是买入今天的股票后所得现金即：-prices[i]
// 那么dp[i][0]应该选所得现金最大的，所以dp[i][0] = max(dp[i - 1][0], -prices[i]);

// 如果第i天不持有股票即dp[i][1]， 也可以由两个状态推出来

// 第i-1天就不持有股票，那么就保持现状，所得现金就是昨天不持有股票的所得现金 即：dp[i - 1][1]
// 第i天卖出股票，所得现金就是按照今天股票价格卖出后所得现金即：prices[i] + dp[i - 1][0]
// 同样dp[i][1]取最大的，dp[i][1] = max(dp[i - 1][1], prices[i] + dp[i - 1][0]);
