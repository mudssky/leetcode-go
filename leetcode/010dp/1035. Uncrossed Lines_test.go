package leetcode

import "testing"

func Test_maxUncrossedLines(t *testing.T) {
	type args struct {
		A []int
		B []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{

		{"test01", args{[]int{1, 4, 2}, []int{1, 2, 4}}, 2},
		{"test02", args{[]int{2, 5, 1, 2, 5}, []int{10, 5, 2, 1, 5, 2}}, 3},
		{"test03", args{[]int{1, 3, 7, 1, 7, 5}, []int{1, 9, 2, 5, 1}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxUncrossedLines(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("maxUncrossedLines() = %v, want %v", got, tt.want)
			}
		})
	}
}
