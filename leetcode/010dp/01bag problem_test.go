package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_zeroOneBag(t *testing.T) {
	type args struct {
		weight    []int
		value     []int
		bagweight int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test01", args{[]int{1, 3, 4}, []int{15, 20, 30}, 4}, 30},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := zeroOneBag(tt.args.weight, tt.args.value, tt.args.bagweight)
			assert.Equal(t, tt.want, got)
		})
	}
}
