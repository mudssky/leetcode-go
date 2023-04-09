package leetcode

import (
	"testing"

	"github.com/mudssky/leetcode-go/structures"
	"github.com/stretchr/testify/assert"
)

func Test_insertIntoBST(t *testing.T) {
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
		{"test01", args{structures.Ints2TreeNode([]int{4, 2, 7, 1, 3}), 5}, []int{4, 2, 7, 1, 3, 5}},
		{"test02", args{structures.Ints2TreeNode([]int{40, 20, 60, 10, 30, 50, 70}), 25}, []int{40, 20, 60, 10, 30, 50, 70, structures.NULL, structures.NULL, 25}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := insertIntoBST(tt.args.root, tt.args.val)
			res := structures.Tree2ints(got)
			assert.Equal(t, tt.want, res)
		})
	}
}
