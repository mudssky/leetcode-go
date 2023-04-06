package leetcode

import (
	"reflect"
	"testing"

	"github.com/mudssky/leetcode-go/structures"
)

func Test_binaryTreePaths(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{"test01", args{structures.Ints2TreeNode([]int{1, 2, 3, structures.NULL, 5})}, []string{"1->2->5", "1->3"}},
		{"test02", args{structures.Ints2TreeNode([]int{1})}, []string{"1"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binaryTreePaths(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("binaryTreePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
