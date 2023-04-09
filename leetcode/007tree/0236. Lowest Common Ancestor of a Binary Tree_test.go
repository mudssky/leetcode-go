package leetcode

import (
	"testing"

	"github.com/mudssky/leetcode-go/structures"
	"github.com/stretchr/testify/assert"
)

func Test_lowestCommonAncestor(t *testing.T) {
	type args struct {
		root *TreeNode
		p    int
		q    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"test01", args{structures.Ints2TreeNode([]int{3, 5, 1, 6, 2, 0, 8, structures.NULL, structures.NULL, 7, 4}), 5, 1}, 3},
		{"test02", args{structures.Ints2TreeNode([]int{3, 5, 1, 6, 2, 0, 8, structures.NULL, structures.NULL, 7, 4}), 5, 4}, 5},
		{"test03", args{structures.Ints2TreeNode([]int{1, 2}), 1, 2}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.args.root.FindNode(tt.args.p)
			q := tt.args.root.FindNode(tt.args.q)
			got := lowestCommonAncestor(tt.args.root, p, q)
			assert.Equal(t, tt.want, got.Val)

		})
	}
}
