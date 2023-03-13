package leetcode

// 「快乐数」定义为：对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和，然后重复这个过程直到这个数变为 1，
// 也可能是 无限循环 但始终变不到 1。如果 可以变为  1，那么这个数就是快乐数。
// 如果 n 是快乐数就返回 True ；不是，则返回 False 。
func isHappy(n int) bool {
	resMap := make(map[int]bool, 0)
	newNum := n
	for {
		newNum = addDigitSquare(newNum)
		if newNum == 1 {
			return true
		}
		if _, ok := resMap[newNum]; ok {
			return false
		}
		resMap[newNum] = true
	}
}

func addDigitSquare(n int) int {
	sum := 0
	num := n
	for num >= 1 {
		digit := num % 10
		num /= 10
		sum += digit * digit
	}
	return sum
}
