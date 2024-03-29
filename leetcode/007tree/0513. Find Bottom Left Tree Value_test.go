package leetcode

import (
	"testing"

	"github.com/mudssky/leetcode-go/structures"
)

func Test_findBottomLeftValue(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"test01", args{structures.Ints2TreeNode([]int{2, 1, 3})}, 1},
		{"test02", args{structures.Ints2TreeNode([]int{1, 2, 3, 4, structures.NULL, 5, 6, structures.NULL, structures.NULL, 7})}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findBottomLeftValue(tt.args.root); got != tt.want {
				t.Errorf("findBottomLeftValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
