package leetcode

import (
	"testing"

	"github.com/mudssky/leetcode-go/structures"
)

func Test_maxDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"test01", args{structures.Ints2TreeNode([]int{3, 9, 20, structures.NULL, structures.NULL, 15, 7})}, 3},
		{"test02", args{structures.Ints2TreeNode([]int{1, structures.NULL, 2})}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
