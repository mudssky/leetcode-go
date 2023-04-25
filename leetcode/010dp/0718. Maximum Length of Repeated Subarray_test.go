package leetcode

import "testing"

func Test_findLength(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{

		{"test01", args{[]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}}, 3},
		{"test02", args{[]int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLength(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
