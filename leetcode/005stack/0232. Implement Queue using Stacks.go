package leetcode

import "github.com/mudssky/leetcode-go/structures"

type MyQueue struct {
	stackIn  *structures.Stack[int]
	stackOut *structures.Stack[int]
}

func Constructor() MyQueue {
	q := new(MyQueue)
	q.stackIn = structures.NewStack[int]()
	q.stackOut = structures.NewStack[int]()
	return *q
}

func (this *MyQueue) Push(x int) {
	this.stackIn.Push(x)
}

func (this *MyQueue) Pop() int {

	if this.stackOut.Empty() {
		// 输出栈为空时，从输入栈全部导入
		for !this.stackIn.Empty() {
			val := this.stackIn.Pop()
			this.stackOut.Push(val)
		}
	}
	return this.stackOut.Pop()
}

func (this *MyQueue) Peek() int {
	if this.stackOut.Empty() {
		for !this.stackIn.Empty() {
			val := this.stackIn.Pop()
			this.stackOut.Push(val)

		}
	}
	return this.stackOut.Peek()
}

func (this *MyQueue) Empty() bool {
	return this.stackIn.Empty() && this.stackOut.Empty()
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
