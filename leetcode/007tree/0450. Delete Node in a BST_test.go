package leetcode

import (
	"testing"

	"github.com/mudssky/leetcode-go/structures"
	"github.com/stretchr/testify/assert"
)

func Test_deleteNode(t *testing.T) {
	type args struct {
		root *TreeNode
		key  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"test01", args{structures.Ints2TreeNode([]int{5, 3, 6, 2, 4, structures.NULL, 7}), 3}, []int{5, 4, 6, 2, structures.NULL, structures.NULL, 7}},
		{"test02", args{structures.Ints2TreeNode([]int{5, 3, 6, 2, 4, structures.NULL, 7}), 0}, []int{5, 3, 6, 2, 4, structures.NULL, 7}},
		{"test03", args{structures.Ints2TreeNode([]int{}), 0}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := deleteNode(tt.args.root, tt.args.key)
			list := structures.Tree2ints(got)
			assert.Equal(t, tt.want, list)
		})
	}
}
