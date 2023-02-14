package util

// 二分查找
// 传入数字，返回找到的下标
func BinarySearch[T Number](arr []T, target T) int {
	high := len(arr) - 1
	low := 0
	for low <= high {
		// 利用位移实现除2，提升性能
		// 先减后加，这样更不容易溢出
		mid := low + (high-low)>>1
		if arr[mid] > target {
			high = mid - 1
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
