package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyQueue(t *testing.T) {
	q := Constructor()
	q.stackIn.Display()
	assert.True(t, q.Empty())
	q.Push(1)
	q.Push(2)
	q.Push(3)
	assert.False(t, q.Empty())
	assert.Equal(t, 1, q.Peek())
	assert.Equal(t, 1, q.Pop())
	assert.Equal(t, 2, q.Peek())
	assert.Equal(t, 2, q.Pop())
	assert.Equal(t, 3, q.Peek())
	assert.Equal(t, 3, q.Pop())
	assert.True(t, q.Empty())
}
