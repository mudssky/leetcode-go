package leetcode

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDecreaseQueue(t *testing.T) {
	// 创建一个新的DecreaseQueue
	q := NewDecreaseQueue[int]()
	// 断言q不是nil
	assert.NotNil(t, q)
	// 断言q.list是空切片
	assert.Empty(t, q.list)
}

func TestBack(t *testing.T) {
	// 创建一个新的DecreaseQueue
	q := NewDecreaseQueue[int]()
	// 向队列中添加一些元素
	q.list = append(q.list, 1, 2, 3)
	// 断言q不是空的
	assert.False(t, q.Empty())
	// 断言q.Back()返回了3
	assert.Equal(t, 3, q.Back())
}

func TestSize(t *testing.T) {
	// 创建一个新的DecreaseQueue
	q := NewDecreaseQueue[int]()
	// 断言q是空的
	assert.True(t, q.Empty())
	// 断言q.Size()返回了0
	assert.Equal(t, 0, q.Size())
	// 向队列中添加一些元素
	q.list = append(q.list, 1, 2, 3)
	// 断言q不是空的
	assert.False(t, q.Empty())
	// 断言q.Size()返回了3
	assert.Equal(t, 3, q.Size())
}

func TestEmpty(t *testing.T) {
	// 创建一个新的DecreaseQueue
	q := NewDecreaseQueue[int]()
	// 断言q是空的
	assert.True(t, q.Empty())
	// 向队列中添加一个元素
	q.list = append(q.list, 1)
	// 断言q不是空的
	assert.False(t, q.Empty())
}

func TestPop(t *testing.T) {
	// 创建一个新的DecreaseQueue
	q := NewDecreaseQueue[int]()
	// 向队列中添加一些元素
	q.list = append(q.list, 1, 2, 3)
	// 断言q不是空的
	assert.False(t, q.Empty())
	// 断言q.Pop()返回了1，并且从队列中移除了它
	q.Pop(1)
	assert.Equal(t, []int{2, 3}, q.list)
	// 断言q.Size()返回了2
	assert.Equal(t, 2, q.Size())
}

func Test_maxSlidingWindow(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"test01", args{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3}, []int{3, 3, 5, 5, 6, 7}},
		{"test01", args{[]int{1}, 1}, []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSlidingWindow(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxSlidingWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
