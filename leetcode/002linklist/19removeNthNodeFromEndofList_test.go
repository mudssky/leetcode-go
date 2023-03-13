package leetcode

import (
	"testing"

	"github.com/mudssky/leetcode-go/structures"
	"github.com/stretchr/testify/assert"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"test1", args{structures.Ints2List([]int{1, 2, 3, 4, 5}), 2}, []int{1, 2, 3, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeNthFromEnd(tt.args.head, tt.args.n)
			list := structures.List2Ints(got)
			assert.Equal(t, tt.want, list)
			// if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			// }
		})
	}
}
