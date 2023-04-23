package leetcode

// 给定一个正整数 n，将其拆分为至少两个正整数的和，并使这些整数的乘积最大化。 返回你可以获得的最大乘积。
func integerBreak(n int) int {
	// dp为拆分n能获得的最大乘积
	dp := make([]int, n+1)
	// 初始化dp，
	dp[2] = 1
	for i := 3; i <= n; i++ {
		for j := 1; j < i-1; j++ {
			// 对比两种拆分方式
			// 一种是j * (i - j) ，另一种是dp[i - j] * dp[j]
			// dp[i-j]指的是拆分两个或者以上的情况，和不进行拆分比较。
			dp[i] = max(dp[i], max((i-j)*j, dp[i-j]*j))
		}
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
