package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fourSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{"test01", args{[]int{1, 0, -1, 0, -2, 2}, 0}, [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}},
		{"test02", args{[]int{2, 2, 2, 2, 2}, 8}, [][]int{{2, 2, 2, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fourSum(tt.args.nums, tt.args.target)
			// t.Log("got", got)
			for _, v := range tt.want {
				assert.Contains(t, got, v)
			}
			assert.Equal(t, len(tt.want), len(got))
		})
	}
}
