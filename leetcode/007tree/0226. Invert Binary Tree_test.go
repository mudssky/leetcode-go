package leetcode

import (
	"testing"

	"github.com/mudssky/leetcode-go/structures"
	"github.com/stretchr/testify/assert"
)

func Test_invertTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"test01", args{structures.Ints2TreeNode([]int{4, 2, 7, 1, 3, 6, 9})}, []int{4, 7, 2, 9, 6, 3, 1}},
		{"test02", args{structures.Ints2TreeNode([]int{2, 1, 3})}, []int{2, 3, 1}},
		{"test03", args{structures.Ints2TreeNode([]int{})}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := invertTree(tt.args.root)
			assert.Equal(t, tt.want, structures.Tree2ints(got))
		})
	}
}
