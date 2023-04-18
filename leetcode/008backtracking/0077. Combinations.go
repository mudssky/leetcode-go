package leetcode

// 给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
func combine(n int, k int) [][]int {
	path, res := make([]int, 0, k), make([][]int, 0)
	var dfs func(n, k, start int)
	// start是开始的数字
	dfs = func(n, k, start int) {
		if len(path) == k { // 说明已经满足了k个数的要求
			tmp := make([]int, k)
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		for i := start; i <= n; i++ { // 从start开始，不往回走，避免出现重复组合
			// 如果元素个数，小于需要满足的组合数，后续没必要往下走了。
			if n-i+1 < k-len(path) { // 剪枝
				break
			}
			path = append(path, i)
			dfs(n, k, i+1)
			path = path[:len(path)-1]
		}
	}
	dfs(n, k, 1)
	return res

}
