package leetcode

// 给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
func combine(n int, k int) [][]int {
	var res [][]int
	var path []int
	var dfs func(n, k, startIndex int)
	dfs = func(n, k, startIndex int) {
		if len(path) == k { // 说明已经满足了k个数的要求
			tmp := make([]int, k)
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		for i := startIndex; i <= n; i++ { // 从start开始，不往回走，避免出现重复组合
			if n-i+1 < k-len(path) { // 剪枝
				break
			}
			path = append(path, i)
			dfs(n, k, i+1)
			path = path[:len(path)-1]
		}
	}
	return res

}
