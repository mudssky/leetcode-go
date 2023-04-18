package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_subsets(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{

		{"test01", args{[]int{1, 2, 3}}, [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}},
		{"test02", args{[]int{0}}, [][]int{{}, {0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := subsets(tt.args.nums)
			assert.Equal(t, len(tt.want), len(got))
			// assert.Equal(t, tt.want, got)
		})
	}
}
