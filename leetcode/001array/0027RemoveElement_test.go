package leetcode

import (
	"sort"
	"testing"

	"github.com/mudssky/leetcode-go/util"
	"github.com/stretchr/testify/assert"
)

// 测试二分查找
func Test0027RemoveElement(t *testing.T) {
	testCases := []struct {
		input1  []int
		input2  int
		expect  int
		expect2 []int
	}{
		{[]int{3, 2, 2, 3}, 3, 2, []int{2, 2}},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5, []int{0, 1, 4, 0, 3}},
	}

	for index, tCases := range testCases {
		res := removeElement(tCases.input1, tCases.input2)
		assert.Equal(t, tCases.expect, res, util.TestErrorMessage(util.TestCaseError{Index: index, Info: tCases}))
		// 先排序再比较，因为这道题是不要求顺序的
		sort.Ints(tCases.expect2)
		sort.Ints(tCases.input1[:res])
		assert.Equal(t, tCases.expect2, tCases.input1[:res], util.TestErrorMessage(util.TestCaseError{Index: index, Info: tCases}))
	}

}
