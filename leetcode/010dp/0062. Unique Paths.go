package leetcode

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	// 初始化dp
	// 左，上两条边都是只有单方向直走1种路线，
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// 到达一个点的路线数量，等于从上面和左面两个位置的路线数量加起来
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
