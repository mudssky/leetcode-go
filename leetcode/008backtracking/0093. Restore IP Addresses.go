package leetcode

import (
	"strconv"
	"strings"
)

// 给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。
func restoreIpAddresses(s string) []string {
	res := []string{}
	path := []string{}

	// pointNum 记录逗号的数量，到4个就是终止条件
	var dfs func(startIndex int)
	dfs = func(startIndex int) {
		// 切出4段，不再继续
		if len(path) == 4 {
			if startIndex >= len(s) {
				str := strings.Join(path, ".")
				res = append(res, str)
			}
			return
		}
		for i := startIndex; i < len(s); i++ {
			// 含有前导0.无效，i>startIndex是排除0.0.0.0这种情况
			if i > startIndex && s[startIndex] == '0' {
				break
			}
			str := s[startIndex : i+1]

			if isValidIPValue(str) {
				path = append(path, str)
				dfs(i + 1)
				path = path[:len(path)-1]
			} else {
				break
				//不满足条件的情况，后续循环也不满足
			}
		}
	}
	dfs(0)
	return res
}
func isValidIPValue(s string) bool {
	num, _ := strconv.Atoi(s)
	if num >= 0 && num <= 255 {
		return true
	}
	return false
}
