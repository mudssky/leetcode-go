package leetcode

import "testing"

func Test_combinationSum4(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test01", args{[]int{1, 2, 3}, 4}, 7},
		{"test02", args{[]int{9}, 3}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum4(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("combinationSum4() = %v, want %v", got, tt.want)
			}
		})
	}
}
