package leetcode

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	// 如果起点出现障碍，直接返回0
	// if obstacleGrid[0][0] == 1 {
	// 	return 0
	// }
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	dp := make([][]int, m)
	// 开辟数组空间
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// 初始化赋值
	// 碰到障碍物就停止往后赋1
	for i := 0; i < m && obstacleGrid[i][0] == 0; i++ {
		dp[i][0] = 1
	}

	for j := 0; j < n && obstacleGrid[0][j] == 0; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// 不碰到障碍物计算，碰到障碍物的就是初始值0
			if obstacleGrid[i][j] != 1 {
				// 到达一个点的路线数量，等于从上面和左面两个位置的路线数量加起来
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}

		}
	}
	return dp[m-1][n-1]
}
