package leetcode

import (
	"testing"

	"github.com/mudssky/leetcode-go/structures"
	"github.com/stretchr/testify/assert"
)

func Test_lowestCommonAncestorBST(t *testing.T) {
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
		{"test01", args{structures.Ints2TreeNode([]int{6, 2, 8, 0, 4, 7, 9, structures.NULL, structures.NULL, 3, 5}), 2, 8}, 6},
		{"test02", args{structures.Ints2TreeNode([]int{6, 2, 8, 0, 4, 7, 9, structures.NULL, structures.NULL, 3, 5}), 2, 4}, 2},
		{"test03", args{structures.Ints2TreeNode([]int{2, 1}), 2, 1}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.args.root.FindNode(tt.args.p)
			q := tt.args.root.FindNode(tt.args.q)
			got := lowestCommonAncestorBST(tt.args.root, p, q)
			assert.Equal(t, tt.want, got.Val)

		})
	}
}
