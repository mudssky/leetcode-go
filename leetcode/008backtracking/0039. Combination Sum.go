package leetcode

import (
	"sort"
)

// 给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
// candidates 中的数字可以无限制重复被选取。
// 所有数字（包括 target）都是正整数。  解集不能包含重复的组合。
// func combinationSum(candidates []int, target int) [][]int {
// 	res := [][]int{}
// 	path := []int{}
// 	// sort.Ints(candidates)
// 	var dfs func(start, sum int)
// 	dfs = func(start, sum int) {
// 		if sum > target {
// 			return
// 		}
// 		if sum == target {
// 			tmp := make([]int, len(path))
// 			copy(tmp, path)
// 			res = append(res, tmp)
// 			return
// 		}
// 		// rest := target - sum
// 		for i := start; i < len(candidates); i++ {

// 			path = append(path, candidates[i])
// 			// sum += candidates[i]
// 			dfs(i, sum+candidates[i])
// 			// sum -= candidates[i]
// 			path = path[:len(path)-1]
// 		}
// 	}
// 	dfs(0, 0)
// 	return res
// }

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	path := []int{}
	sort.Ints(candidates)
	var dfs func(start, need int)
	dfs = func(start, need int) {
		// fmt.Println(start, need)
		if need == 0 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := start; i < len(candidates); i++ {
			// 剪枝
			if candidates[i] > need {
				break
			}
			path = append(path, candidates[i])
			// sum += candidates[i]
			dfs(i, need-candidates[i])
			// sum -= candidates[i]
			path = path[:len(path)-1]
		}
	}
	dfs(0, target)
	return res
}
