package leetcode

import (
	"reflect"
	"testing"
)

func Test_generateMatrix(t *testing.T) {

	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{

		{"4", args{4}, [][]int{{1, 2, 3, 4}, {12, 13, 14, 5}, {11, 16, 15, 6}, {10, 9, 8, 7}}},
		{"3", args{3}, [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}},
		{"1", args{1}, [][]int{{1}}},
	}
	for _, tt := range tests {
		// fmt.Println(generateMatrix(tt.args.n))
		t.Run(tt.name, func(t *testing.T) {
			if got := generateMatrix(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
