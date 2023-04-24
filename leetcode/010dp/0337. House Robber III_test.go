package leetcode

import (
	"testing"

	"github.com/mudssky/leetcode-go/structures"
)

func Test_rob3(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test01", args{structures.Ints2TreeNode([]int{
			3, 2, 3, structures.NULL, 3, structures.NULL, 1,
		})}, 7},
		{"test02", args{structures.Ints2TreeNode([]int{
			3, 4, 5, 1, 3, structures.NULL, 1,
		})}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob3(tt.args.root); got != tt.want {
				t.Errorf("rob3() = %v, want %v", got, tt.want)
			}
		})
	}
}
