package leetcode

import (
	. "github.com/mudssky/leetcode-go/structures"
)

// 用链表做栈效率太低了
func isBracketMatch2(left, right rune) bool {
	bracketMap := map[rune]rune{'(': ')', '{': '}', '[': ']'}
	if r, ok := bracketMap[left]; ok && r == right {
		return true
	}
	return false
}

// 经典的栈用于括号匹配
func isValid2(s string) bool {
	stack := NewStack[rune]()
	for _, v := range s {
		if stack.Empty() {
			stack.Push(v)
		} else {
			p := stack.Peek()
			if isBracketMatch2(p, v) {
				stack.Pop()
			} else {
				stack.Push(v)
			}
		}
	}
	return stack.Empty()
}

func isBracketMatch(left, right byte) bool {
	// bracketMap := map[byte]byte{'(': ')', '{': '}', '[': ']'}
	// if r, ok := bracketMap[left]; ok && r == right {
	// 	return true
	// }
	if (left == '(' && right == ')') || (left == '[' && right == ']') || (left == '{' && right == '}') {
		return true
	}
	return false
}

// 直接用数组模拟栈
func isValid(s string) bool {
	sLen := len(s)

	// 排除一些情况
	if sLen%2 != 0 {
		return false
	}
	if sLen == 0 {
		return true
	}

	stack := make([]byte, sLen)
	stackTop := -1
	for i := 0; i < sLen; i++ {
		if stackTop == -1 {
			stackTop++
			stack[stackTop] = s[i]
		} else {
			top := stack[stackTop]
			if isBracketMatch(top, s[i]) {
				stackTop--
			} else {
				stackTop++
				stack[stackTop] = s[i]
			}
		}
	}
	return stackTop == -1
}
