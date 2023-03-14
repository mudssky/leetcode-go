package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseStr(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"test01", args{"abcdefg", 2}, "bacdfeg"},
		{"test02", args{"abcd", 2}, "bacd"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseStr(tt.args.s, tt.args.k)
			fmt.Println(got)
			assert.Equal(t, tt.want, got)
		})
	}
}
