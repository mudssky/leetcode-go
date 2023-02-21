package leetcode

// 给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的 连续 子数组，并返回其长度。如果不存在符合条件的子数组，返回 0。
// 示例：
// 输入：s = 7, nums = [2,3,1,2,4,3] 输出：2 解释：子数组 [4,3] 是该条件下的长度最小的子数组。

// 超时了
// func minSubArrayLen(target int, nums []int) int {
// 	length := len(nums)
// 	minLength := length + 1
// 	for i := 0; i < length; i++ {
// 		if nums[i] >= target {
// 			return 1
// 		}
// 		sum := nums[i]
// 		for j := i + 1; j < length; j++ {
// 			sum += nums[j]
// 			if sum >= target {
// 				currentLen := j - i + 1
// 				if minLength > currentLen {
// 					minLength = currentLen

// 				}
// 				break

// 			}
// 		}
// 	}
// 	if minLength == length+1 {
// 		return 0
// 	}
// 	return minLength
// }

func minSubArrayLen(target int, nums []int) int {
	length := len(nums)
	minLength := length + 1
	left := 0
	sum := 0
	for right := 0; right < length; right++ {
		sum += nums[right]
		for sum >= target {
			subLength := right - left + 1
			if subLength < minLength {
				minLength = subLength
			}
			sum -= nums[left]
			left++
		}

	}
	if minLength == length+1 {
		return 0
	}
	return minLength
}
