package leetcode

import (
	"testing"

	"github.com/mudssky/leetcode-go/structures"
	"github.com/stretchr/testify/assert"
)

func Test_convertBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"test01", args{structures.Ints2TreeNode([]int{4, 1, 6, 0, 2, 5, 7, structures.NULL, structures.NULL, structures.NULL, 3, structures.NULL, structures.NULL, structures.NULL, 8})}, []int{30, 36, 21, 36, 35, 26, 15, structures.NULL, structures.NULL, structures.NULL, 33, structures.NULL, structures.NULL, structures.NULL, 8}},

		{"test02", args{structures.Ints2TreeNode([]int{0, structures.NULL, 1})}, []int{1, structures.NULL, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := convertBST(tt.args.root)
			list := structures.Tree2ints(got)
			assert.Equal(t, tt.want, list)
		})
	}
}
