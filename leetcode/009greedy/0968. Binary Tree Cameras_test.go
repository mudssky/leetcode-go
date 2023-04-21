package leetcode

import (
	"testing"

	"github.com/mudssky/leetcode-go/structures"
)

func Test_minCameraCover(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test01", args{structures.Ints2TreeNode([]int{0, 0, structures.NULL, 0, 0})}, 1},
		{"test02", args{structures.Ints2TreeNode([]int{0, 0, structures.NULL, 0, structures.NULL, 0, structures.NULL, structures.NULL, 0})}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCameraCover(tt.args.root); got != tt.want {
				t.Errorf("minCameraCover() = %v, want %v", got, tt.want)
			}
		})
	}
}
