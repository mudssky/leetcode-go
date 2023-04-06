package leetcode

import (
	"testing"

	"github.com/mudssky/leetcode-go/structures"
)

func Test_isBalanced(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"test01", args{structures.Ints2TreeNode([]int{3, 9, 20, structures.NULL, structures.NULL, 15, 7})}, true},
		{"test02", args{structures.Ints2TreeNode([]int{1, 2, 2, 3, 3, structures.NULL, structures.NULL, 4, 4})}, false},
		{"test03", args{structures.Ints2TreeNode([]int{})}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBalanced(tt.args.root); got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
