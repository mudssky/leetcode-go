package leetcode

// 斐波那契数，通常用 F(n) 表示，形成的序列称为 斐波那契数列 。该数列由 0 和 1 开始，后面的每一项数字都是前面两项数字的和。
// 也就是： F(0) = 0，F(1) = 1 F(n) = F(n - 1) + F(n - 2)，其中 n > 1 给你n ，请计算 F(n) 。

// 动态规划
func fib(n int) int {
	// 输入限制最大为30，所以定义最大30的数组够了
	var dp [31]int
	dp[0] = 0
	dp[1] = 1
	if n < 2 {
		return dp[n]
	}
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-2] + dp[i-1]
	}
	return dp[n]
}
