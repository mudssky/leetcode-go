package leetcode

import (
	"reflect"
	"testing"

	"github.com/mudssky/leetcode-go/structures"
)

func Test_removeElements(t *testing.T) {
	type args struct {
		list []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{1, 2, 3, 4, 5}, 1}, []int{2, 3, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			linklist := removeElements(structures.Ints2List(tt.args.list), tt.args.val)
			list := structures.List2Ints(linklist)

			if got := list; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
