package leetcode

import (
	"testing"

	"github.com/mudssky/leetcode-go/structures"
	"github.com/stretchr/testify/assert"
)

func Test_constructMaximumBinaryTree(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"test01", args{[]int{3, 2, 1, 6, 0, 5}}, []int{6, 3, 5, structures.NULL, 2, 0, structures.NULL, structures.NULL, 1}},
		{"test02", args{[]int{3, 2, 1}}, []int{3, structures.NULL, 2, structures.NULL, 1}}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := constructMaximumBinaryTree(tt.args.nums)
			assert.Equal(t, tt.want, structures.Tree2ints(got))

		})
	}
}
