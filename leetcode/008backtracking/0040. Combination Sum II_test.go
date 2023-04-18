package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_combinationSum2(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"test01", args{[]int{10, 1, 2, 7, 6, 1, 5}, 8}, [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}}},
		{"test02", args{[]int{2, 5, 2, 1, 2}, 5}, [][]int{{1, 2, 2}, {5}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combinationSum2(tt.args.candidates, tt.args.target)
			assert.Equal(t, len(tt.want), len(got))
			assert.Equal(t, tt.want, got)
		})
	}
}
