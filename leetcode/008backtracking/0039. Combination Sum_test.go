package leetcode

import (
	"reflect"
	"testing"
)

func Test_combinationSum(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{

		{"test01", args{[]int{2, 3, 6, 7}, 7}, [][]int{{2, 2, 3}, {7}}},
		{"test02", args{[]int{2, 3, 5}, 8}, [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}},
		{"test03", args{[]int{2}, 1}, [][]int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum(tt.args.candidates, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinationSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
