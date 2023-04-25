package leetcode

func isSubsequence(s string, t string) bool {
	// dp[i][j] 表示以下标i为结尾的字符串s，和以下标j为结尾的字符串t，相同子序列的长度为dp[i][j]
	dp := make([][]int, len(s)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(t)+1)
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[len(s)][len(t)] == len(s)
}
