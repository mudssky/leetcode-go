package leetcode

// 给定一个整型数组, 你的任务是找到所有该数组的递增子序列，递增子序列的长度至少是2。
// 给定数组的长度不会超过15。
// 数组中的整数范围是 [-100,100]。
// 给定数组中可能包含重复数字，相等的数字应该被视为递增的一种情况。
func findSubsequences(nums []int) [][]int {
	res := [][]int{}
	path := []int{}
	var dfs func(startIndex int)
	dfs = func(startIndex int) {
		if len(path) >= 2 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
		}
		// used数组，用于对同层的元素去重
		// used := make(map[int]bool, len(nums))
		// 使用数组做used查询，因为题目的数据范围是-100 到100，所以我们映射到201大小的数组上
		used := make([]bool, 201)
		// var used [len(nums)]bool
		for i := startIndex; i < len(nums); i++ {

			// used里面能找到，说明这个数字已经用过了 去重
			// 不是递增的情况，也进入下一次循环
			if used[nums[i]+100] || len(path) > 0 && nums[i] < path[len(path)-1] {
				continue
			}
			path = append(path, nums[i])
			used[nums[i]+100] = true
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return res
}
