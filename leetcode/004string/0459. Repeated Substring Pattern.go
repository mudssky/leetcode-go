package leetcode

// func repeatedSubstringPattern(s string) bool {
// 	sLen := len(s)
// 	doubleS := s + s
// 	middleS := doubleS[1 : len(doubleS)-1]
// 	for i := 0; i < len(middleS)-sLen+1; i++ {
// 		if middleS[i:i+len(s)] == s {
// 			return true
// 		}
// 	}
// 	return false
// }

// 可以用之前的strstr函数
// func repeatedSubstringPattern(s string) bool {
// 	// sLen := len(s)
// 	doubleS := s + s
// 	middleS := doubleS[1 : len(doubleS)-1]
// 	index := strStr(middleS, s)
// 	// for i := 0; i < len(middleS)-sLen+1; i++ {
// 	// 	if middleS[i:i+len(s)] == s {
// 	// 		return true
// 	// 	}
// 	// }
// 	if index == -1 {
// 		return false
// 	}
// 	return true
// }

// 使用kmp算法
func repeatedSubstringPattern(s string) bool {
	n := len(s)
	if n == 0 {
		return false
	}
	j := 0
	next := make([]int, n)
	next[0] = j
	for i := 1; i < n; i++ {
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		next[i] = j
	}
	// next[n-1]  最长相同前后缀的长度
	if next[n-1] != 0 && n%(n-next[n-1]) == 0 {
		return true
	}
	return false
}
