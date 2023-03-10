package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyLinkedList(t *testing.T) {
	/**
	 * Your MyLinkedList object will be instantiated and called as such:
	 * linkedList := Constructor();
	 * param_1 := linkedList.Get(index);
	 * linkedList.AddAtHead(val);
	 * linkedList.AddAtTail(val);
	 * linkedList.AddAtIndex(index,val);
	 * linkedList.DeleteAtIndex(index);
	 */
	// linkedList := Constructor()
	// linkedList.AddAtHead(1)
	// linkedList.AddAtTail(3)
	// linkedList.AddAtIndex(1, 2)

	// assert.Equal(t, 2, linkedList.Get(1))
	// linkedList.DeleteAtIndex(1)
	// linkedList.DisplayList()
	// assert.Equal(t, 3, linkedList.Get(1))

	linkedList := Constructor()
	linkedList.AddAtHead(7)
	linkedList.AddAtHead(2)
	linkedList.AddAtHead(1)
	linkedList.AddAtIndex(3, 0)

	// linkedList.DisplayList()
	linkedList.DeleteAtIndex(2)
	assert.Equal(t, 0, linkedList.Get(2))
	// linkedList.DisplayList()
	linkedList.AddAtHead(6)

	linkedList.AddAtTail(4)
	// linkedList.DisplayList()
	assert.Equal(t, 4, linkedList.Get(4))

	linkedList2 := Constructor()

	linkedList2.AddAtIndex(1, 0)
	assert.Equal(t, -1, linkedList2.Get(0))

	linkedList3 := Constructor()
	linkedList3.AddAtHead(1)
	linkedList3.AddAtTail(3)
	linkedList3.DisplayList()
	linkedList3.AddAtIndex(1, 2)
	linkedList3.DisplayList()
	assert.Equal(t, 2, linkedList3.Get(1))
	linkedList3.DeleteAtIndex(1)
	assert.Equal(t, 3, linkedList3.Get(1))

}
