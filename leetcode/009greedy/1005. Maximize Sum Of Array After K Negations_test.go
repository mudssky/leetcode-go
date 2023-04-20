package leetcode

import "testing"

func Test_largestSumAfterKNegations(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{

		{"test01", args{[]int{4, 2, 3}, 1}, 5},
		{"test02", args{[]int{3, -1, 0, 2}, 3}, 6},
		{"test03", args{[]int{2, -3, -1, 5, -4}, 2}, 13},
		{"test04", args{[]int{-8, 3, -5, -3, -5, -2}, 6}, 22},
		{"test05", args{[]int{-4, -2, -3}, 4}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestSumAfterKNegations(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("largestSumAfterKNegations() = %v, want %v", got, tt.want)
			}
		})
	}
}
