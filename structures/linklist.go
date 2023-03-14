package structures

type LinkList[T comparable] interface {
	Len() int
	// InsertAfter() *Element[T]
	PushBack() *Element[T]
	PushFront() *Element[T]
	Remove(*Element[T]) T
}

type Element[T comparable] struct {
	Val  T
	Next *Element[T]
	Prev *Element[T]
}

// 采用双向环形链表+哑节点，这样在头部和尾部插入都比较方便。
type threadUnsafeLinkList[T comparable] struct {
	dummy  *Element[T]
	length int
}

// 工厂方法
func NewThreadUnsafeLinkList[T comparable]() *threadUnsafeLinkList[T] {
	// var zero T
	// rear := &Element[T]{
	// 	Val:  zero,
	// 	Next: nil,
	// 	Prev: nil,
	// }
	// rear.Next = rear
	// rear.Prev = rear
	// return &threadUnsafeLinkList[T]{dummy: rear, length: 0}
	return new(threadUnsafeLinkList[T]).Init()
}

func (l *threadUnsafeLinkList[T]) Init() *threadUnsafeLinkList[T] {
	var zero T
	rear := &Element[T]{
		Val:  zero,
		Next: nil,
		Prev: nil,
	}
	rear.Next = rear
	rear.Prev = rear
	l.dummy = rear
	l.length = 0
	return l
}

func (l *threadUnsafeLinkList[T]) Len() int {
	return l.length
}

func (l *threadUnsafeLinkList[T]) PushBack(val T) *Element[T] {
	newNode := &Element[T]{
		Val:  val,
		Next: l.dummy,
		Prev: l.dummy.Prev,
	}
	// dummy哑节点的Prev就是链表结尾
	l.dummy.Prev.Next = newNode
	l.dummy.Prev = newNode
	l.length++
	return l.dummy.Prev
}

func (l *threadUnsafeLinkList[T]) PushFront(val T) *Element[T] {
	newNode := &Element[T]{
		Val:  val,
		Next: l.dummy.Next,
		Prev: l.dummy,
	}
	l.dummy.Next.Prev = newNode
	l.dummy.Next = newNode
	l.length++
	return l.dummy.Next
}

func (l *threadUnsafeLinkList[T]) Remove(node *Element[T]) T {
	var res T
	if node == nil {
		return res
	}
	head := l.dummy.Next
	for head != l.dummy {
		if head == node {
			// 找到相等的节点，进行删除操作
			head.Prev.Next = head.Next
			head.Next.Prev = head.Prev
			l.length--
			return head.Val

		}
		head = head.Next
	}

	return res
}

// 转换列表为切片
func (l *threadUnsafeLinkList[T]) Slice() []T {
	var res []T
	l.ForEach(func(value T, index int) {
		res = append(res, value)
	})
	return res
}
func (l *threadUnsafeLinkList[T]) ForEach(iteratee func(value T, index int)) {
	head := l.dummy.Next
	index := 0
	for head != l.dummy {
		iteratee(head.Val, index)
		head = head.Next
	}

}
