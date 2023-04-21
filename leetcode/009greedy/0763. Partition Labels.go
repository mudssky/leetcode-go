package leetcode

// 字符串 S 由小写字母组成。我们要把这个字符串划分为尽可能多的片段，同一字母最多出现在一个片段中。返回一个表示每个字符串片段的长度的列表。

func partitionLabels(s string) []int {
	var res []int
	var marks [26]int
	size, left, right := len(s), 0, 0
	// 记录字母出现的最远下标
	for i := 0; i < size; i++ {
		marks[s[i]-'a'] = i
	}
	for i := 0; i < size; i++ {
		// 字符串的右边界
		right = max(right, marks[s[i]-'a'])
		// 如果字母的下标恰好是最远下标，就是分割点。
		if i == right {
			res = append(res, right-left+1)
			left = i + 1
		}
	}
	return res
}
