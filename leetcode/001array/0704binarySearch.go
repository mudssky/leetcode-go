package leetcode

// 给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。
// 704
func search(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	// 已经排除了空数组的情况，此时high=-1，进入不了循环
	for low <= high {
		mid := low + (high-low)>>1
		if target > nums[mid] {
			low = mid + 1
		} else if target < nums[mid] {
			high = mid - 1
		} else {
			return mid
		}
	}
	return -1
	// return util.BinarySearch(nums, target)
}
