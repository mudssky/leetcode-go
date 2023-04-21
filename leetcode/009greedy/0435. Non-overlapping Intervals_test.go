package leetcode

import "testing"

func Test_eraseOverlapIntervals(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test01", args{[][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}}, 1},
		{"test02", args{[][]int{{1, 2}, {1, 2}, {1, 2}}}, 2},
		{"test03", args{[][]int{{1, 2}, {2, 3}}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := eraseOverlapIntervals(tt.args.intervals); got != tt.want {
				t.Errorf("eraseOverlapIntervals() = %v, want %v", got, tt.want)
			}
		})
	}
}
