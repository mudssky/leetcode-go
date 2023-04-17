package leetcode

import (
	"testing"

	"github.com/mudssky/leetcode-go/structures"
	"github.com/stretchr/testify/assert"
)

func Test_trimBST(t *testing.T) {
	type args struct {
		root *TreeNode
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"test01", args{structures.Ints2TreeNode([]int{1, 0, 2}), 1, 2}, []int{1, structures.NULL, 2}},
		{"test02", args{structures.Ints2TreeNode([]int{3, 0, 4, structures.NULL, 2, structures.NULL, structures.NULL, 1}), 1, 3}, []int{3, 2, structures.NULL, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := trimBST(tt.args.root, tt.args.low, tt.args.high)
			list := structures.Tree2ints(got)
			assert.Equal(t, tt.want, list)
		})
	}
}
