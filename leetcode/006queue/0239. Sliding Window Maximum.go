package leetcode

import "github.com/mudssky/leetcode-go/constraints"

// 单调递减队列，头部始终是最大值
type DecreaseQueue[T constraints.Ordered] struct {
	list []T
}

func NewDecreaseQueue[T constraints.Ordered]() *DecreaseQueue[T] {
	return &DecreaseQueue[T]{
		list: []T{},
	}
}
func (q *DecreaseQueue[T]) Front() T {
	return q.list[0]
}
func (q *DecreaseQueue[T]) Back() T {
	return q.list[q.Size()-1]
}
func (q *DecreaseQueue[T]) Size() int {
	return len(q.list)
}
func (q *DecreaseQueue[T]) Empty() bool {
	return q.Size() == 0
}

// 单调队列，pop传入一个值比较
func (q *DecreaseQueue[T]) Pop(val T) {
	//  如果val值等于单调队列的出口元素front，那么队列弹出元素，否则不用任何操作
	// 也就是随着窗口划过，队列的元素不在窗口范围内，所以要移除
	if !q.Empty() && val == q.Front() {
		q.list = q.list[1:]
	}
}

func (q *DecreaseQueue[T]) Push(val T) {
	// 维护列表的排序，如果新增的值比列表尾部大，则把尾部的元素移除，
	// 直到队列为空，或者val正好小于尾部的值
	for !q.Empty() && val > q.Back() {
		q.list = q.list[:q.Size()-1]
	}
	q.list = append(q.list, val)
}

func maxSlidingWindow(nums []int, k int) []int {
	numsLen := len(nums)
	dq := NewDecreaseQueue[int]()
	// 先把窗口内的元素放入单调队列
	for i := 0; i < k; i++ {
		dq.Push(nums[i])
	}
	res := make([]int, 0)
	// 记录前k个元素的最大值
	res = append(res, dq.Front())
	for i := k; i < numsLen; i++ {
		// 滑动窗口移除最前面的元素
		dq.Pop(nums[i-k])
		// 滑动窗口添加最后面的元素
		dq.Push(nums[i])
		// 记录最大值
		res = append(res, dq.Front())
	}
	return res
}
