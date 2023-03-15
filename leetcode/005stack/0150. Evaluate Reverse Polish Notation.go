package leetcode

import "strconv"

// 经典逆波兰表达式求值

func isOperator(s string) bool {
	return s == "+" || s == "-" || s == "*" || s == "/"
}

// func calculateStr(n1s, n2s, op string) string {
// 	n1, _ := strconv.Atoi(n1s)
// 	n2, _ := strconv.Atoi(n2s)
// 	res := -1
// 	switch op {
// 	case "+":
// 		res = n1 + n2
// 	case "-":
// 		res = n1 - n2
// 	case "/":
// 		res = n1 / n2
// 	case "*":
// 		res = n1 * n2
// 	}
// 	return strconv.Itoa(res)
// }
// func evalRPN(tokens []string) int {

// 	tLen := len(tokens)
// 	stack := make([]string, tLen)
// 	stackTop := -1
// 	for i := 0; i < tLen; i++ {
// 		if isOperator(tokens[i]) {
// 			res := calculate(stack[stackTop-1], stack[stackTop], tokens[i])
// 			stackTop--
// 			stack[stackTop] = res
// 		} else {
// 			stackTop++
// 			stack[stackTop] = tokens[i]
// 		}
// 	}
// 	res, _ := strconv.Atoi(stack[stackTop])
// 	return res
// }

func calculate(n1, n2 int, op string) int {
	res := -1
	switch op {
	case "+":
		res = n1 + n2
	case "-":
		res = n1 - n2
	case "/":
		res = n1 / n2
	case "*":
		res = n1 * n2
	}
	return res
}

// 优化一下字符串转换的耗时
func evalRPN(tokens []string) int {

	tLen := len(tokens)
	stack := make([]int, tLen)
	stackTop := -1
	for i := 0; i < tLen; i++ {
		if isOperator(tokens[i]) {
			res := calculate(stack[stackTop-1], stack[stackTop], tokens[i])
			stackTop--
			stack[stackTop] = res
		} else {
			stackTop++
			num, _ := strconv.Atoi(tokens[i])
			stack[stackTop] = num
		}

	}
	return stack[0]
}
