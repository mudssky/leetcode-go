package leetcode

// 找出所有相加之和为 n 的 k 个数的组合。组合中只允许含有 1 - 9 的正整数，并且每种组合中不存在重复的数字。
func combinationSum3(k int, n int) [][]int {
	res, path := make([][]int, 0), make([]int, 0, k)
	var dfs func(k, n, start, sum int)
	dfs = func(k, n, start, sum int) {
		if len(path) == k {
			if sum == n {
				// 拷贝一份，添加到结果
				tmp := make([]int, k)
				copy(tmp, path)
				res = append(res, tmp)

			}
			return
		}
		for i := start; i <= 9; i++ {
			// 剪枝，总和超过n，和组合数目不够的情况，不用继续了。
			if sum+i > n || 9-i+1 < k-len(path) {
				break
			}
			path = append(path, i)
			dfs(k, n, i+1, sum+i)
			path = path[:len(path)-1]
		}

	}
	dfs(k, n, 1, 0)
	return res
}
