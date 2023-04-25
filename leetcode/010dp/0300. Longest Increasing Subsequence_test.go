package leetcode

import "testing"

func Test_lengthOfLIS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test01", args{[]int{10, 9, 2, 5, 3, 7, 101, 18}}, 4},
		{"test02", args{[]int{0, 1, 0, 3, 2, 3}}, 4},
		{"test03", args{[]int{7, 7, 7, 7, 7, 7, 7}}, 1},
		{"test04", args{[]int{1, 3, 6, 7, 9, 4, 10, 5, 6}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLIS(tt.args.nums); got != tt.want {
				t.Errorf("lengthOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}
