package leetcode

import (
	"testing"

	"github.com/mudssky/leetcode-go/structures"
)

func Test_isValidBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"test01", args{structures.Ints2TreeNode([]int{2, 1, 3})}, true},
		{"test02", args{structures.Ints2TreeNode([]int{5, 1, 4, structures.NULL, structures.NULL, 3, 6})}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
