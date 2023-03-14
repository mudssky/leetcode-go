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

// func (this *MyStack) Push(x int)  {

// }

// func (this *MyStack) Pop() int {

// }

// func (this *MyStack) Top() int {

// }

// func (this *MyStack) Empty() bool {

// }

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
