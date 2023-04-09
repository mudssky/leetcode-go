package leetcode

import (
	"testing"

	"github.com/mudssky/leetcode-go/structures"
)

func Test_getMinimumDifference(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"test01", args{structures.Ints2TreeNode([]int{4, 2, 6, 1, 3})}, 1},
		{"test02", args{structures.Ints2TreeNode([]int{1, 0, 48, structures.NULL, structures.NULL, 12, 49})}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMinimumDifference(tt.args.root); got != tt.want {
				t.Errorf("getMinimumDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
