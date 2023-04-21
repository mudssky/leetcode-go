package leetcode

import "strconv"

// 给定一个非负整数 N，找出小于或等于 N 的最大的整数，同时这个整数需要满足其各个位数上的数字是单调递增。
// （当且仅当每个相邻位数上的数字 x 和 y 满足 x <= y 时，我们称这个整数是单调递增的。）

// 先写一个暴力求解的算法
// func monotoneIncreasingDigits(n int) int {
// 	for i := n; i >= 0; i-- {
// 		if isMonotoneIncrease(i) {
// 			return i
// 		}
// 	}
// 	return -1
// }

// 检测一个数是否每一位都是单调递增
func isMonotoneIncrease(n int) bool {
	if n < 10 {
		return true
	}
	// 先取最后一位
	pre := n % 10
	n = n / 10
	// 循环取下一位（更高位），如果cur大于pre说明不满足单调递增
	for n > 0 {
		cur := n % 10
		if cur > pre {
			return false
		}
		pre = cur
		n = n / 10
	}
	return true
}

// 用贪心优化
// 遍历一个数字，从后向前，如果numstr[i-1]>numstr[i],此时我们让numstr[i-1]-1 ,numstr[i]变成9
// 举个例子，数字：332，从前向后遍历的话，那么就把变成了329，此时2又小于了第一位的3了，真正的结果应该是299。
// 也就是说，最后除了最开头和一开始递增的部分，其他的数字都会变成9
func monotoneIncreasingDigits(n int) int {
	// 还是转成字符串比较更直观
	numstr := []byte(strconv.Itoa(n))
	// flag用于判断9开始的位置
	flag := len(numstr)
	for i := len(numstr) - 1; i > 0; i-- {
		if numstr[i-1] > numstr[i] {
			flag = i
			numstr[i-1] -= 1
		}
	}
	for i := flag; i < len(numstr); i++ {
		numstr[i] = '9'
	}

	res, _ := strconv.Atoi(string(numstr))
	return res
}
