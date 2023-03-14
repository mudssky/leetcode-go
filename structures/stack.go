package structures

import "fmt"

type Stacker[T comparable] interface {
	Push(val T)
	Pop() T
	Peek() T
	Empty() bool
	Size() int
}

type Stack[T comparable] struct {
	list threadUnsafeLinkList[T]
}

func (s *Stack[T]) Init() {
	s.list = *NewThreadUnsafeLinkList[T]()
}

func NewStack[T comparable]() *Stack[T] {
	s := new(Stack[T])
	s.Init()
	return s
}

func (s *Stack[T]) Push(val T) {
	s.list.PushBack(val)
}

func (s *Stack[T]) Pop() T {
	var res T
	if !s.Empty() {
		return s.list.Remove(s.list.dummy.Prev)
	}
	return res
}

func (s *Stack[T]) Empty() bool {
	return s.list.length == 0
}
func (s *Stack[T]) Size() int {
	return s.list.length
}
func (s *Stack[T]) Peek() T {
	var res T
	if !s.Empty() {
		res = s.list.dummy.Prev.Val
	}
	return res
}

func (s *Stack[T]) Display() {
	fmt.Println(s.list.Slice())
}
