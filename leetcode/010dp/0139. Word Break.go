package leetcode

// 给定一个非空字符串 s 和一个包含非空单词的列表 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。

// func wordBreak(s string, wordDict []string) bool {
// 	// 装满背包s的前几位字符的方式有几种
// 	dp := make([]int, len(s)+1)
// 	dp[0] = 1
// 	for i := 0; i <= len(s); i++ { // 背包
// 		for j := 0; j < len(wordDict); j++ { // 物品
// 			// 背包大于字符串长度，并且到i为止，可以找到相等的字符串
// 			if i >= len(wordDict[j]) && wordDict[j] == s[i-len(wordDict[j]):i] {
// 				dp[i] += dp[i-len(wordDict[j])]
// 			}
// 		}
// 	}

// 	return dp[len(s)] > 0
// }

func wordBreak(s string, wordDict []string) bool {
	wordDictSet := make(map[string]bool)
	for _, w := range wordDict {
		wordDictSet[w] = true
	}
	// dp[i] : 字符串长度为i的话，dp[i]为true，表示可以拆分为一个或多个在字典中出现的单词。
	dp := make([]bool, len(s)+1)
	dp[0] = true
	// 本题求得是排列数
	// 外层遍历背包，内层遍历物品
	for i := 1; i < len(s); i++ {
		// 如果确定dp[j] 是true，且 [j, i] 这个区间的子串出现在字典里，那么dp[i]一定是true。（j < i ）。
		for j := 0; j < i; j++ {
			if dp[j] && wordDictSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
