package leetcode

import (
	"reflect"
	"testing"
)

func Test_combinationSum3(t *testing.T) {
	type args struct {
		k int
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"test01", args{3, 7}, [][]int{{1, 2, 4}}},
		{"test02", args{3, 9}, [][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}}},
		{"test03", args{4, 1}, [][]int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := combinationSum3(tt.args.k, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinationSum3() = %v, want %v", got, tt.want)
			}
		})
	}
}
