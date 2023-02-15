package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	name     string
	input    []int
	expected []int
}{
	{
		name:     "test1",
		input:    []int{4, 2, 3, 1, 5},
		expected: []int{1, 2, 3, 4, 5},
	},
	{
		name:     "test2",
		input:    []int{10, 8, 6, 4, 2},
		expected: []int{2, 4, 6, 8, 10},
	},
	{
		name:     "test3",
		input:    []int{1, 2, 3, 4, 5},
		expected: []int{1, 2, 3, 4, 5},
	},
}

func TestBubbleSort(t *testing.T) {
	for index, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			BubbleSort(tc.input)
			assert.Equal(t, tc.input, tc.expected, TestErrorMessage(TestCaseError{Index: index, Info: tc}))
			// if !reflect.DeepEqual(tc.input, tc.expected) {
			// 	t.Errorf("expected %v, but got %v", tc.expected, tc.input)
			// }
		})
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// 生成一个随机数切片
		arr := RandomInts(1000, 1000)
		BubbleSort(arr)
	}
}

func TestInsertSort(t *testing.T) {

	for index, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			InsertSort(tc.input)
			assert.Equal(t, tc.input, tc.expected, TestErrorMessage(TestCaseError{Index: index, Info: tc}))
		})
	}
}
