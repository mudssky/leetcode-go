package util

import (
	"fmt"

	"github.com/mudssky/leetcode-go/constraints"
)

// 转换整数为字符串
func Itoa[T constraints.Integer](num T) string {
	// Sprintf使用反射效率比较低，以后应该还有优化空间
	return fmt.Sprintf("%d", num)
}
