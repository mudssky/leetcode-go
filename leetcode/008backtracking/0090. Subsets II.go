package leetcode

import "sort"

// 给定一个可能包含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
// 说明：解集不能包含重复的子集。
func subsetsWithDup(nums []int) [][]int {
	res := [][]int{}
	path := []int{}
	// 排序，这样之后可以跳过
	sort.Ints(nums)
	var dfs func(startIndex int)
	dfs = func(startIndex int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)

		for i := startIndex; i < len(nums); i++ {
			// 跳过重复项
			if i > startIndex && nums[i] == nums[i-1] {
				continue
			}
			path = append(path, nums[i])
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return res
}
