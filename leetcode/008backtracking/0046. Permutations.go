package leetcode

// 给定一个 没有重复 数字的序列，返回其所有可能的全排列。

func permute(nums []int) [][]int {
	path := []int{}
	res := [][]int{}
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
			// path中已经使用过的元素，跳过，因为一个排列里一个元素只能使用一次。
			if used[i] {
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
