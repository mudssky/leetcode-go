package leetcode

import (
	"testing"

	"github.com/mudssky/leetcode-go/util"
	"github.com/stretchr/testify/assert"
)

// 测试二分查找
func Test0977SquaresofaSortedArray(t *testing.T) {
	testCases := []struct {
		input  []int
		expect []int
	}{
		{[]int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}},
		{[]int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121}},
	}
	for index, tcase := range testCases {
		res := sortedSquares(tcase.input)
		assert.Equal(t, tcase.expect, res, util.TestErrorMessage(util.TestCaseError{Index: index, Info: tcase}))
	}
}
