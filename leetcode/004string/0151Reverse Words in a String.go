package leetcode

import (
	"strings"
)

func reverseWords(s string) string {
	sList := strings.Fields(s)
	reverseList(sList)
	return strings.TrimSpace(strings.Join(sList, " "))
}

func reverseList[T comparable](list []T) {
	listLen := len(list)
	for i := 0; i < listLen/2; i++ {
		list[i], list[listLen-1-i] = list[listLen-1-i], list[i]
	}
}

// 发现我以前写的更简单
// js用多了以后，习惯用函数来处理了。

// func reverseWords(s string) string {
// 	strList := strings.Split(s, " ")
// 	resList := []string{}
// 	for i := len(strList) - 1; i >= 0; i-- {
// 		if len(strList[i]) > 0 && strList[i][0] != ' ' {
// 			resList = append(resList, strList[i])
// 		}
// 	}

// 	return strings.Join(resList, " ")
// }
