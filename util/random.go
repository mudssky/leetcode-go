package util

import (
	"math/rand"
	"time"
)

func RandomInt(n int) int {
	rand.Seed(time.Now().UnixNano())

	// 生成一个范围在 [0, 100) 之间的随机整数
	return rand.Intn(100)
}
