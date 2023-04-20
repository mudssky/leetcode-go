package leetcode

import "math"

// func maxSubArray(nums []int) int {
// 	maxSum := nums[0]
// 	for i := 1; i < len(nums); i++ {
// 		if nums[i]+nums[i-1] > nums[i] {
// 			nums[i] += nums[i-1]
// 		}
// 		if nums[i] > maxSum {
// 			maxSum = nums[i]
// 		}
// 	}
// 	return maxSum
// }

func maxSubArray(nums []int) int {
	res := math.MinInt
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum > res {
			res = sum
		}
		// 变成负数时重置，因为和负数相加值只会变小
		if sum <= 0 {
			sum = 0
		}
	}
	return res
}
