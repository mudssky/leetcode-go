package leetcode

import (
	"testing"

	"github.com/mudssky/leetcode-go/structures"
	"github.com/stretchr/testify/assert"
)

func Test_inorderTraversal(t *testing.T) {
	type args struct {
		treelist []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"test01", args{[]int{1, structures.NULL, 2, 3}}, []int{1, 3, 2}},
		{"test02", args{[]int{}}, []int{}},
		{"test03", args{[]int{1}}, []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inorderTraversal(structures.Ints2TreeNode(tt.args.treelist))
			assert.Equal(t, tt.want, got)
		})
	}

}
