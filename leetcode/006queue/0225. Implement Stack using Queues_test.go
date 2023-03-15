package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyStack(t *testing.T) {
	// 创建一个新的MyStack对象
	s := Constructor()
	// 断言s是否为空
	assert.True(t, s.Empty(), "s should be empty")
	// 向s中压入两个元素
	s.Push(1)
	s.Push(2)
	// 断言s不为空
	assert.False(t, s.Empty(), "s should not be empty")
	// 断言s的顶部元素是2
	assert.Equal(t, 2, s.Top(), "top of s should be 2")
	// 弹出s的顶部元素并断言其值是2
	assert.Equal(t, 2, s.Pop(), "pop from s should return 2")
	// 断言s的顶部元素是1
	assert.Equal(t, 1, s.Top(), "top of s should be 1")
}
