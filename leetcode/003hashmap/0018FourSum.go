package leetcode

import "sort"

// 给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c
// + d 的值与 target 相等？找出所有满足条件且不重复的四元组。

func fourSum(nums []int, target int) [][]int {
	numsLen := len(nums)
	if numsLen < 4 {
		return nil
	}
	sort.Ints(nums)
	res := [][]int{}
	for i := 0; i < numsLen-3; i++ {
		n1 := nums[i]
		if i > 0 && n1 == nums[i-1] {
			continue
		}
		for j := i + 1; j < numsLen-2; j++ {
			n2 := nums[j]
			if j > i+1 && n2 == nums[j-1] { // 对nums[j]去重
				continue
			}
			l, r := j+1, numsLen-1
			for l < r {
				n3, n4 := nums[l], nums[r]
				sumOf4 := n1 + n2 + n3 + n4
				if sumOf4 == target {
					res = append(res, []int{n1, n2, n3, n4})
					// 去重逻辑应该放在找到一个三元组之后，对b 和 c去重
					for l < r && nums[l] == n3 {
						l++
					}
					for l < r && nums[r] == n4 {
						r--
					}
				} else if sumOf4 < target {
					l++
				} else {
					r--
				}
			}
		}
	}
	return res
}
