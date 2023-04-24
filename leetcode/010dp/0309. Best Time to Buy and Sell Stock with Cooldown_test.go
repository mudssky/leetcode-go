package leetcode

import "testing"

func Test_maxProfit5(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test01", args{[]int{1, 2, 3, 0, 2}}, 3},
		{"test02", args{[]int{1}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit5(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit5() = %v, want %v", got, tt.want)
			}
		})
	}
}
