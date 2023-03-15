package leetcode

import "github.com/mudssky/leetcode-go/structures"

type MyStack struct {
	queue *structures.Queue[int]
}

func Constructor() MyStack {
	s := new(MyStack)
	s.queue = structures.NewQueue[int]()
	return *s
}

func (this *MyStack) Push(x int) {
	this.queue.Push(x)
}

func (this *MyStack) Pop() int {
	// return this.queue.Pop()
	size := this.queue.Size()
	for size > 1 {
		size--
		this.queue.Push(this.queue.Pop())
	}
	return this.queue.Pop()
}

func (this *MyStack) Top() int {
	return this.queue.Back()
}

func (this *MyStack) Empty() bool {
	return this.queue.Empty()
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
