package leetcode

import (
	"testing"

	"github.com/mudssky/leetcode-go/structures"
	"github.com/stretchr/testify/assert"
)

func Test_isSymmetric(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"test01", args{structures.Ints2TreeNode([]int{1, 2, 2, 3, 4, 4, 3})}, true},
		{"test02", args{structures.Ints2TreeNode([]int{1, 2, 3, structures.NULL, 3, structures.NULL, 3})}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isSymmetric(tt.args.root)
			assert.Equal(t, tt.want, got)
		})
	}
}
