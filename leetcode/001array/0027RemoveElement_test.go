package leetcode

import (
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
		// input := make([]int, len(tCases.input1))
		// copy(input, tCases.input1)
		// fmt.Println("tt", tCases)
		res := removeElement(tCases.input1, tCases.input2)
		// fmt.Println("input", input)
		assert.Equal(t, tCases.expect, res, util.TestErrorMessage(util.TestCaseError{Index: index, Info: tCases}))
		assert.Equal(t, tCases.expect2, tCases.input1[:res], util.TestErrorMessage(util.TestCaseError{Index: index, Info: tCases}))
	}
}
