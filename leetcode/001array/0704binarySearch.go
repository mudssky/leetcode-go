package leetcode

// 给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。
// 704
// 时间复杂度是O(logn) 空间复杂度O(1),几乎没有使用额外空间
func search(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	// 已经排除了空数组的情况，此时high=-1，进入不了循环
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
		// 当 low 和 high 接近时，切换到线性查找,避免二分查找的额外开销
		// if high-low <= 16 {
		// 	for i := low; i <= high; i++ {
		// 		if nums[i] == target {
		// 			return i
		// 		}
		// 	}
		// 	return -1
		// }
	}
	return -1
	// return util.BinarySearch(nums, target)
}
