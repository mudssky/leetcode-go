package leetcode

import (
	"testing"

	"github.com/mudssky/leetcode-go/structures"
	"github.com/stretchr/testify/assert"
)

func Test_buildTree(t *testing.T) {
	type args struct {
		inorder   []int
		postorder []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"test01", args{[]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}}, []int{3, 9, 20, structures.NULL, structures.NULL, 15, 7}},
		{"test02", args{[]int{-1}, []int{-1}}, []int{-1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildTree(tt.args.inorder, tt.args.postorder)
			assert.Equal(t, tt.want, structures.Tree2ints(got))
		})
	}
}
