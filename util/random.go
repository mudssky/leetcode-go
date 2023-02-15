package util

import (
	"math/rand"
	"time"
)

// 生成一个范围在 [0, n) 之间的随机整数
func RandomInt(n int) int {
	rand.Seed(time.Now().UnixNano())

	// 生成一个范围在 [0, n) 之间的随机整数
	return rand.Intn(n)
}

// 生成一个范围在 [0, n) 之间的随机整数组成的长度count的列表
func RandomInts(num int, count int) (res []int) {
	res = make([]int, count)
	for i := 0; i < count; i++ {
		res[i] = RandomInt(num)
	}
	return
}
