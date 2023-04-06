package leetcode

import (
	"testing"

	"github.com/mudssky/leetcode-go/structures"
	"github.com/stretchr/testify/assert"
)

func Test_searchBST(t *testing.T) {
	type args struct {
		root *TreeNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"test01", args{structures.Ints2TreeNode([]int{4, 2, 7, 1, 3}), 2}, []int{2, 1, 3}},
		{"test02", args{structures.Ints2TreeNode([]int{4, 2, 7, 1, 3}), 5}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := searchBST(tt.args.root, tt.args.val)
			assert.Equal(t, tt.want, structures.Tree2ints(got))
		})
	}
}
