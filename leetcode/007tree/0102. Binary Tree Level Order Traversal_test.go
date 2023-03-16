package leetcode

import (
	"testing"

	"github.com/mudssky/leetcode-go/structures"
	"github.com/stretchr/testify/assert"
)

func Test_levelOrder(t *testing.T) {
	type args struct {
		root []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{"test01", args{[]int{3, 9, 20, structures.NULL, structures.NULL, 15, 7}}, [][]int{{3}, {9, 20}, {15, 7}}},
		{"test02", args{[]int{1}}, [][]int{{1}}},
		{"test03", args{[]int{}}, [][]int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := levelOrder(structures.Ints2TreeNode(tt.args.root))
			assert.Equal(t, tt.want, got)
		})
	}
}
