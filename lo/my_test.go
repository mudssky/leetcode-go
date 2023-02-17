package lo

import (
	"reflect"
	"testing"
)

func TestConcat(t *testing.T) {
	type args struct {
		collection []int
		values     []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{1, 2, 3}, []int{4, 5}}, []int{1, 2, 3, 4, 5}},
		{"test2", args{[]int{2, 3}, []int{4, 5}}, []int{2, 3, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Concat(tt.args.collection, tt.args.values...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Concat() = %v, want %v", got, tt.want)
			}
		})
	}
}
