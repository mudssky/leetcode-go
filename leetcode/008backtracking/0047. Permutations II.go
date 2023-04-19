package leetcode

import "sort"

// 这道题目和46.全排列 的区别在与给定一个可包含重复数字的序列，要返回所有不重复的全排列。
func permuteUnique(nums []int) [][]int {
	path := []int{}
	res := [][]int{}
	// 先对数组排序，方便后面去重
	sort.Ints(nums)
	// 全局使用一个used数组，记录path中使用过的坐标
	used := make(map[int]bool, len(nums))
	var dfs func()
	dfs = func() {
		if len(path) >= len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		for i := 0; i < len(nums); i++ {
			// 1.used[i]  path中已经使用过的元素，跳过，因为一个排列里一个元素只能使用一次。
			// 2.跳过重复的元素, 因为used[i-1]==false, 说明是同一树层，加上这个判断其实是筛选树枝和树层、、、
			// 这里是判断树层中出现了重复，进行去重
			// used[i-1]=true,则说明是同一树枝
			if used[i] || i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			used[i] = true
			path = append(path, nums[i])

			dfs()
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	dfs()
	return res
}
