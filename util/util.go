package util

// 二分查找
// 传入数字，返回找到的下标
// 找不到的情况返回-1

func BinarySearch[T Number](arr []T, target T) int {
	low := 0
	high := len(arr) - 1
	// 已经排除了空数组的情况，此时high=-1，进入不了循环
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
