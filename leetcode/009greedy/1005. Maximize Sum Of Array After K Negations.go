package leetcode

import (
	"math"
	"sort"
)

// 贪心算法
func largestSumAfterKNegations2(nums []int, k int) int {
	// 1.按照绝对值大小递减排序
	sort.Slice(nums, func(i, j int) bool {
		return math.Abs(float64(nums[i])) > math.Abs(float64(nums[j]))
	})

	// 2.从小到大遍历负数，将其转变为证书
	for i := 0; i < len(nums); i++ {
		if k > 0 && nums[i] < 0 {
			nums[i] = -nums[i]
			k--
		}
	}

	// 3.如果还有剩余的k为奇数，则转换一个最小的正数，消费掉。
	// 如果为偶数，则是对一个数反复使用浪费掉
	if k%2 == 1 {
		nums[len(nums)-1] = -nums[len(nums)-1]
	}

	// 4.把结果加起来
	result := 0
	for i := 0; i < len(nums); i++ {
		result += nums[i]
	}
	return result
}

// 有一个特殊情况，就是负数和正数分界的地方，这里如果k次数还没用完，就要对比上一个负数的回到负数，
// 和把正数变成负数哪个负数更大，也就是哪个绝对值更小
func largestSumAfterKNegations(nums []int, k int) int {
	// 没必要按照绝对值排序
	sort.Ints(nums)
	res := 0
	negIndex := 0
	// 一次循环可以完成操作
	for i := 0; i < len(nums); i++ {
		if k > 0 {
			// 负数的情况，取反
			if nums[i] < 0 {
				nums[i] = -nums[i]
				k--
				negIndex = i
			} else {
				// 碰到正数和0
				// 只在k为奇数次时，执行一次，把其他次数重复使用浪费掉
				// 这一次需要在第一个正数，和最后一个
				if k%2 == 1 {
					if nums[negIndex] < nums[i] {
						// nums[negIndex] = -nums[negIndex]
						// 这个情况，res需要额外执行减掉,因为已经遍历过了negIndex
						res -= nums[negIndex] + nums[negIndex]
					} else {
						nums[i] = -nums[i]
					}
				}
				// 用掉所有的k
				k = 0
			}

		}

		// 处理完以后，加到结果
		res += nums[i]
	}
	// 循环完后k仍不为0的情况
	// 减去末尾的负数
	if k != 0 {
		res -= nums[len(nums)-1] * 2
	}
	return res
}

func minAbs(a, b int) int {
	var pa int
	var pb int
	if a < 0 {
		pa = -a
	}
	if b < 0 {
		pb = -b
	}
	if pa > pb {
		return b
	}
	return a
}
