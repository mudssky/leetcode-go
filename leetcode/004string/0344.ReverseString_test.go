package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
		{"test01", args{[]byte{'h', 'e', 'l', 'l', 'o'}}, []byte{'o', 'l', 'l', 'e', 'h'}},
		{"test01", args{[]byte{'H', 'a', 'n', 'n', 'a', 'h'}}, []byte{'h', 'a', 'n', 'n', 'a', 'H'}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseString(tt.args.s)
			assert.Equal(t, string(tt.want), string(tt.args.s))
		})
	}
}
