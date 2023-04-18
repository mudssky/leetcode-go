package leetcode

// 给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。
func partition(s string) [][]string {
	path := []string{}
	res := [][]string{}

	var dfs func(startIndex int)

	dfs = func(startIndex int) {
		// 如果起始位置等于s的大小，说明已经找到一组分割方案了。
		if startIndex >= len(s) {
			tmp := make([]string, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		for i := startIndex; i < len(s); i++ {
			str := s[startIndex : i+1]
			if isPalindrome(str) {
				path = append(path, str)
				dfs(i + 1)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return res
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
