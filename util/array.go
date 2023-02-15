package util

// 这个文件里面是数组相关操作

// 原地修改输入数组，，移除元素，返回移除元素后的长度
func RemoveInPlace[T comparable](arr []T, target T) int {
	length := len(arr)
	res := 0
	for i := 0; i < length; i++ {
		if arr[i] != target {
			arr[res] = arr[i]
			res++
		}
	}
	return res
}
