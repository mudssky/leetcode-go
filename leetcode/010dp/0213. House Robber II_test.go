package leetcode

import "testing"

func Test_rob2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test01", args{[]int{2, 3, 2}}, 3},
		{"test02", args{[]int{1, 2, 3, 1}}, 4},
		{"test03", args{[]int{1, 2, 3}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob2(tt.args.nums); got != tt.want {
				t.Errorf("rob2() = %v, want %v", got, tt.want)
			}
		})
	}
}
