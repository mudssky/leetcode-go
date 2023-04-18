package leetcode

func subsets(nums []int) [][]int {
	res := [][]int{}
	path := []int{}
	var dfs func(startIndex int)
	dfs = func(startIndex int) {
		// 第一次是空集，之后每次进入，就把path存入结果集（而不是遍历到叶节点才是结果集）
		// 因为求子集就是遍历整颗树，所以不需要剪枝
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)

		for i := startIndex; i < len(nums); i++ {
			path = append(path, nums[i])
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return res
}
