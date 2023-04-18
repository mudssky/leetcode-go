package leetcode

// 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合
func letterCombinations(digits string) []string {
	path := []byte{}
	res := []string{}
	if digits == "" {
		return res
	}
	var lMap = map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	var dfs func(start int)
	dfs = func(start int) {
		if len(path) == len(digits) {
			res = append(res, string(path))
			return
		}
		str := lMap[digits[start]]
		for i := 0; i < len(str); i++ {
			path = append(path, str[i])
			dfs(start + 1)
			// 回溯，清理现场，以便继续下一个循环
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return res
}
