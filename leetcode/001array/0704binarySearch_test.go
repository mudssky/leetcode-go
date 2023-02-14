package leetcode

import (
	"testing"

	"github.com/mudssky/leetcode-go/util"
	"github.com/stretchr/testify/assert"
)

// 测试二分查找
func Test0704BinarySearch(t *testing.T) {
	testCases := []struct {
		input1 []int
		input2 int
		expect int
	}{
		{[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{[]int{-1, 0, 3, 5, 9, 12}, 2, -1},
	}
	for index, tcase := range testCases {
		res := search(tcase.input1, tcase.input2)
		assert.Equal(t, tcase.expect, res, util.TestErrorMessage(util.TestCaseError{Index: index, Info: tcase}))
	}
}
