package leetcode

func isAnagram(s string, t string) bool {
	// 字节数组，足够存储26个字母的信息
	charArr := make([]int, 26)
	// 遍历第一个字符串，
	for _, c := range s {
		charArr[c-'a']++
	}
	// 如果第二个字符串是第一个字符串重新组合，那么相减以后，数组的每一位应该正好是0
	for _, c := range t {
		index := c - 'a'
		charArr[index]--
	}
	for _, item := range charArr {
		if item != 0 {
			return false
		}
	}
	return true
}
