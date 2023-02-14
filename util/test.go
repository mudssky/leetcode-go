package util

import "fmt"

type TestCaseError struct {
	Index int //测试用例的下标
	Info  any // 测试用例相关信息，包括 输入，和期望输出
}

// 拼接测试输出
func TestErrorMessage(info TestCaseError) string {
	return fmt.Sprintf("case %d fails,%#v", info.Index, info.Info)
}
