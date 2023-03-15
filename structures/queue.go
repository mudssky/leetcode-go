package structures

type Queuer[T comparable] interface {
	Push(val T)
	Pop() T
	Peek() T
	Empty() bool
	Size() int
	Back() T
}

type Queue[T comparable] struct {
	list threadUnsafeLinkList[T]
}

func (q *Queue[T]) Init() {
	q.list = *NewThreadUnsafeLinkList[T]()
}

func NewQueue[T comparable]() *Queue[T] {
	q := new(Queue[T])
	q.Init()
	return q
}
func (q *Queue[T]) Push(val T) {
	q.list.PushBack(val)
}

func (q *Queue[T]) Pop() T {
	var res T
	if !q.Empty() {
		return q.list.Remove(q.list.dummy.Next)
	}
	return res
}
func (q *Queue[T]) Empty() bool {
	return q.list.length == 0
}

func (q *Queue[T]) Peek() T {
	var res T
	if !q.Empty() {
		res = q.list.dummy.Next.Val
	}
	return res
}

func (q *Queue[T]) Size() int {
	return q.list.length
}

func (q *Queue[T]) Back() T {
	var res T
	if !q.Empty() {
		return q.list.dummy.Prev.Val
	}
	return res
}
