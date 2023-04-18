package leetcode

import "sort"

// 给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
// candidates 中的每个数字在每个组合中只能使用一次。

// func combinationSum2(candidates []int, target int) [][]int {
// 	res := [][]int{}
// 	path := []int{}
// 	// 通过树层去重，需要先对数组排序
// 	sort.Ints(candidates)

// 	// used数组，用于去重
// 	used := make([]bool, len(candidates))

// 	var dfs func(start, need int)
// 	dfs = func(start, need int) {
// 		if need == 0 {
// 			tmp := make([]int, len(path))
// 			copy(tmp, path)
// 			res = append(res, tmp)
// 		}
// 		if need < 0 {
// 			return
// 		}
// 		for i := start; i < len(candidates); i++ {
// 			if candidates[i] > need { // 剪枝，提前返回
// 				break
// 			}

// 			// used[i - 1] == true，说明同一树枝candidates[i - 1]使用过
// 			// used[i - 1] == false，说明同一树层candidates[i - 1]使用过
// 			// 这里我们要去重的是同一树层上使用过的。
// 			if i > 0 && candidates[i] == candidates[i-1] && !used[i-1] {
// 				continue
// 			}
// 			path = append(path, candidates[i])
// 			used[i] = true
// 			dfs(i+1, need-candidates[i])
// 			used[i] = false
// 			path = path[:len(path)-1]

// 		}
// 	}
// 	dfs(0, target)
// 	return res
// }

func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	path := []int{}
	// 通过树层去重，需要先对数组排序
	sort.Ints(candidates)

	var dfs func(start, need int)
	dfs = func(start, need int) {
		if need == 0 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
		}
		if need < 0 {
			return
		}
		for i := start; i < len(candidates); i++ {
			// 可以不用used数组，反正只要把重复的序列跳过就行了。
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}

			path = append(path, candidates[i])

			dfs(i+1, need-candidates[i])

			path = path[:len(path)-1]

		}
	}
	dfs(0, target)
	return res
}
