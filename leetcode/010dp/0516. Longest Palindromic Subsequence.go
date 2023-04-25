package leetcode

// 给定一个字符串 s ，找到其中最长的回文子序列，并返回该序列的长度。可以假设 s 的最大长度为 1000 。

// 示例 1: 输入: "bbbab" 输出: 4 一个可能的最长回文子序列为 "bbbb"。

// 示例 2: 输入:"cbbd" 输出: 2 一个可能的最长回文子序列为 "bb"。

func longestPalindromeSubseq(s string) int {
	// dp[i][j]：字符串s在[i, j]范围内最长的回文子序列的长度为dp[i][j]。
	size := len(s)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	dp := make([][]int, size)
	for i := 0; i < size; i++ {
		dp[i] = make([]int, size)
		dp[i][i] = 1
	}
	for i := size - 1; i >= 0; i-- {
		for j := i + 1; j < size; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i+1][j])
			}
		}
	}
	return dp[0][size-1]

}
