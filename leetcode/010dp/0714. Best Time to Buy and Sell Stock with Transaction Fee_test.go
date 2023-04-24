package leetcode

import "testing"

func Test_maxProfit6(t *testing.T) {
	type args struct {
		prices []int
		fee    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test01", args{[]int{1, 3, 2, 8, 4, 9}, 2}, 8},
		{"test02", args{[]int{1, 3, 7, 5, 10, 3}, 3}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit6(tt.args.prices, tt.args.fee); got != tt.want {
				t.Errorf("maxProfit6() = %v, want %v", got, tt.want)
			}
		})
	}
}
