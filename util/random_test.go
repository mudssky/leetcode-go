package util

import "testing"

func TestRandomInt(t *testing.T) {
	n := 100
	for i := 0; i < 100; i++ {
		res := RandomInt(n)
		if res < 0 || res >= n {
			t.Errorf("RandomInt(%d) = %d; want >= 0 and < %d", n, res, n)
		}
	}
}

func TestRandomInts(t *testing.T) {
	num := 100
	count := 10
	res := RandomInts(num, count)
	if len(res) != count {
		t.Errorf("RandomInts(%d, %d) returned %d elements; want %d", num, count, len(res), count)
	}
	for _, v := range res {
		if v < 0 || v >= num {
			t.Errorf("RandomInts(%d, %d) contains %d; want >= 0 and < %d", num, count, v, num)
		}
	}
}
