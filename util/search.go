package util

import (
	constraints "github.com/mudssky/leetcode-go/constraints"
)

// 这个文件放搜索算法相关

// 二分查找
// 传入升序排序的数字列表，返回找到的下标
// 找不到的情况返回-1

func BinarySearch[T constraints.Ordered](arr []T, target T) int {
	low := 0
	high := len(arr) - 1
	// 已经排除了空数组的情况，此时high=-1，进入不了循环
	for low <= high {
		// 利用位移实现除2，提升性能
		// 先减后加，这样更不容易溢出
		mid := low + ((high - low) >> 1)
		if arr[mid] > target {
			high = mid - 1
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			return mid
		}
		// 当 low 和 high 接近时，切换到线性查找,避免二分查找的额外开销
		// 经过benchmark测试发现，并没有优化，反而稍微慢了一点，有优化的情况提升也不明显。
		// if high-low <= 16 {
		// 	for i := low; i <= high; i++ {
		// 		if arr[i] == target {
		// 			return i
		// 		}
		// 	}
		// 	return -1
		// }
	}
	return -1
}

// 线性查找
// 无论是否排序都可以查找,传入可以==比较的元素组成的数组
func LinearSearch[T comparable](arr []T, target T) int {
	for i := range arr {
		if arr[i] == target {
			return i
		}
	}
	return -1
}

// 二分查找性能的对照组
func binarySearchCompare[T constraints.Ordered](arr []T, target T) int {
	low := 0
	high := len(arr) - 1
	// 已经排除了空数组的情况，此时high=-1，进入不了循环
	for low <= high {
		// 利用位移实现除2，提升性能
		// 先减后加，这样更不容易溢出
		mid := low + ((high - low) >> 1)
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
