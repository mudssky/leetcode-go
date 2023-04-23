package leetcode

// 有n件物品和一个最多能背重量为w 的背包。第i件物品的重量是weight[i]，得到的价值是value[i] 。每件物品只能用一次，求解将哪些物品装入背包里物品价值总和最大。
func zeroOneBag(weight, value []int, bagweight int) int {
	// dp 是背包负重和物品重量的二维数组，
	// dp[i][j]的含义：从下标为[0-i]的物品里任意取，放进容量为j的背包，价值总和最大是多少。
	dp := make([][]int, len(weight))
	for i := range dp {
		dp[i] = make([]int, bagweight+1)
	}
	// 初始化dp
	// 背包负重为0的情况，价值一定为0
	// 因为数组初始值就是0，所以不用手动赋值了。

	// 背包负重比编号为0第一个物品大的情况，此时i=0的情况，一定是编号0物品的价值
	for j := bagweight; j > weight[0]; j-- {
		dp[0][j] = value[0]
	}
	// 其他状态可以由状态转移方程推出
	for i := 1; i < len(weight); i++ {
		for j := 0; j <= bagweight; j++ {
			// 背包放不下的情况下，和不放价值相同
			if j < weight[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				// 能放入的情况下，需要和不放的情况比较，选取最大的。
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
			}
		}
	}
	return dp[len(weight)-1][bagweight]

}
