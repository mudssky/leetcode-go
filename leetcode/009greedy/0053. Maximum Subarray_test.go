package leetcode

import "testing"

func Test_maxSubArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"test01", args{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}}, 6},
		{"test02", args{[]int{1}}, 1},
		{"test03", args{[]int{5, 4, -1, 7, 8}}, 23},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray(tt.args.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
