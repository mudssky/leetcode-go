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
		{"test01", args{[]int{1, 3, 4}, []int{15, 20, 30}, 4}, 35},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := zeroOneBag(tt.args.weight, tt.args.value, tt.args.bagweight)
			assert.Equal(t, tt.want, got)

		})
	}
}

func Test_zeroOneBag2(t *testing.T) {
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
		{"test01", args{[]int{1, 3, 4}, []int{15, 20, 30}, 4}, 35},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := zeroOneBag2(tt.args.weight, tt.args.value, tt.args.bagweight)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_complete01Bag(t *testing.T) {
	type args struct {
		weight    []int
		value     []int
		bagWeight int
	}
	tests := []struct {
		name string
		args args
		want int
	}{

		{"test01", args{[]int{1, 3, 4}, []int{15, 20, 30}, 4}, 60},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := complete01Bag(tt.args.weight, tt.args.value, tt.args.bagWeight); got != tt.want {
				t.Errorf("complete01Bag() = %v, want %v", got, tt.want)
			}
		})
	}
}
