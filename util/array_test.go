package util

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveElement(t *testing.T) {
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
		res := RemoveInPlace(tCases.input1, tCases.input2)
		assert.Equal(t, tCases.expect, res, TestErrorMessage(TestCaseError{Index: index, Info: tCases}))
		// 先排序再比较，因为这道题是不要求顺序的
		sort.Ints(tCases.expect2)
		sort.Ints(tCases.input1[:res])
		assert.Equal(t, tCases.expect2, tCases.input1[:res], TestErrorMessage(TestCaseError{Index: index, Info: tCases}))
	}

}
