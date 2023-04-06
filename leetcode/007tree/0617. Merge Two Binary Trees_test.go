package leetcode

import (
	"testing"

	"github.com/mudssky/leetcode-go/structures"
	"github.com/stretchr/testify/assert"
)

func Test_mergeTrees(t *testing.T) {
	type args struct {
		root1 *TreeNode
		root2 *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"test01", args{structures.Ints2TreeNode([]int{1, 3, 2, 5}), structures.Ints2TreeNode([]int{2, 1, 3, structures.NULL, 4, structures.NULL, 7})}, []int{3, 4, 5, 5, 4, structures.NULL, 7}},
		{"test02", args{structures.Ints2TreeNode([]int{1}), structures.Ints2TreeNode([]int{1, 2})}, []int{2, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeTrees(tt.args.root1, tt.args.root2)
			assert.Equal(t, tt.want, structures.Tree2ints(got))
		})
	}
}
