package structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	s := NewStack[int]()
	assert.True(t, s.Empty())
	s.Push(1)
	s.Push(2)
	s.Push(3)
	assert.False(t, s.Empty())
	assert.Equal(t, 3, s.Peek())
	assert.Equal(t, 3, s.Pop())
	assert.Equal(t, 2, s.Peek())
	assert.Equal(t, 2, s.Pop())
	assert.Equal(t, 1, s.Peek())
	assert.Equal(t, 1, s.Pop())
	assert.True(t, s.Empty())
}
