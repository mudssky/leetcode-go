package leetcode

// func removeDuplicates(s string) string {
// 	sLen := len(s)
// 	stack := make([]byte, sLen)
// 	stackTop := -1
// 	for i := 0; i < sLen; i++ {
// 		if stackTop == -1 {
// 			stackTop++
// 			stack[stackTop] = s[i]
// 		} else {
// 			top := stack[stackTop]
// 			if top == s[i] {
// 				stackTop--
// 			} else {
// 				stackTop++
// 				stack[stackTop] = s[i]
// 			}
// 		}
// 	}
// 	return string(stack[:stackTop+1])
// }

// 之前有一个问题是空间申请比较多，虽然性能提升了。。。
func removeDuplicates(s string) string {
	var stack []byte
	for i := 0; i < len(s); i++ {
		// 栈不空 且 与栈顶元素不等
		if len(stack) > 0 && stack[len(stack)-1] == s[i] {
			// 弹出栈顶元素 并 忽略当前元素(s[i])
			stack = stack[:len(stack)-1]
		} else {
			// 入栈
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}
