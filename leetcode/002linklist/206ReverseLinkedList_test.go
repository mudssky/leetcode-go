package leetcode

import (
	"reflect"
	"testing"

	"github.com/mudssky/leetcode-go/structures"
	"github.com/stretchr/testify/assert"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"test01", args{structures.Ints2List([]int{1, 2, 3, 4})}, []int{4, 3, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				ints := structures.List2Ints(got)
				// fmt.Println(ints)
				// t.Errorf("reverseList() = %v, want %v", ints, tt.want)
				assert.Equal(t, tt.want, ints)
			}
		})
	}
}
