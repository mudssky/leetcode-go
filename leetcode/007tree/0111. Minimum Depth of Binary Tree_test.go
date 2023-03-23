package leetcode

import (
	"testing"

	"github.com/mudssky/leetcode-go/structures"
)

func Test_minDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"test01", args{structures.Ints2TreeNode([]int{3, 9, 20, structures.NULL, structures.NULL, 15, 7})}, 2},
		{"test02", args{structures.Ints2TreeNode([]int{2, structures.NULL, 3, structures.NULL, 4, structures.NULL, 5, structures.NULL, 6})}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDepth(tt.args.root); got != tt.want {
				t.Errorf("minDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
