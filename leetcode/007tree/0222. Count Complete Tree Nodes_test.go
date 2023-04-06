package leetcode

import (
	"testing"

	"github.com/mudssky/leetcode-go/structures"
)

func Test_countNodes(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{

		{"test01", args{structures.Ints2TreeNode([]int{1, 2, 3, 4, 5, 6})}, 6},
		{"test02", args{structures.Ints2TreeNode([]int{})}, 0},
		{"test03", args{structures.Ints2TreeNode([]int{1})}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countNodes(tt.args.root); got != tt.want {
				t.Errorf("countNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
