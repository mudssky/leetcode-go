package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{"test01", args{[]int{-1, 0, 1, 2, -1, -4}}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{"test02", args{[]int{0, 1, 1}}, [][]int{}},
		{"test03", args{[]int{0, 0, 0}}, [][]int{{0, 0, 0}}},
		{"test04", args{[]int{0, 0, 0, 0}}, [][]int{{0, 0, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := threeSum(tt.args.nums)

			for _, v := range tt.want {
				assert.Contains(t, got, v)
			}
			assert.Equal(t, len(tt.want), len(got))
		})
	}
}
