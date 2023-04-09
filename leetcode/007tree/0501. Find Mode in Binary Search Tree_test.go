package leetcode

import (
	"reflect"
	"testing"

	"github.com/mudssky/leetcode-go/structures"
)

func Test_findMode(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"test01", args{structures.Ints2TreeNode([]int{1, structures.NULL, 2, 2})}, []int{2}},
		{"test02", args{structures.Ints2TreeNode([]int{0})}, []int{0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMode(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMode() = %v, want %v", got, tt.want)
			}
		})
	}
}
