package mapset

type Set[T comparable] interface {
	// Add adds an element to the set. Returns whether
	// the item was added.
	Add(val T) bool

	// Contains returns whether the given items
	// are all in the set.
	Contains(val ...T) bool
	// Remove removes a single element from the set.
	Remove(i T)
}

type threadUnsafeSet[T comparable] map[T]struct{}

func newThreadUnsafeSet[T comparable]() threadUnsafeSet[T] {
	return make(threadUnsafeSet[T])
}

func (s threadUnsafeSet[T]) Add(v T) bool {
	prevLen := len(s)
	s[v] = struct{}{}
	return prevLen != len(s)
}

func (s threadUnsafeSet[T]) Remove(v T) {
	delete(s, v)
}

func (s threadUnsafeSet[T]) Contains(v ...T) bool {
	for _, val := range v {
		if _, ok := s[val]; !ok {
			return false
		}
	}
	return true
}

// 写个算法不需要线程安全
// 基于go的map,只读是线程安全的,如果要写也线程安全,需要加一个读写锁...
func NewThreadUnsafeSet[T comparable](vals ...T) Set[T] {
	s := newThreadUnsafeSet[T]()
	for _, item := range vals {
		s.Add(item)
	}
	return s
}
