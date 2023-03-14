package structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkList(t *testing.T) {
	l := NewThreadUnsafeLinkList[int]()
	assert.Equal(t, 0, l.Len())
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	assert.Equal(t, 3, l.Len())
	assert.Equal(t, []int{1, 2, 3}, l.Slice())
	e := l.PushFront(0)
	assert.Equal(t, []int{0, 1, 2, 3}, l.Slice())
	v := l.Remove(e)
	assert.Equal(t, 0, v)
	assert.Equal(t, []int{1, 2, 3}, l.Slice())
}
