package leetcode

import "sort"

// 使用排序函数实现
func topKFrequent(nums []int, k int) []int {
	countMap := make(map[int]int)
	for _, v := range nums {
		countMap[v]++
	}
	ans := []int{}
	for k := range countMap {
		ans = append(ans, k)
	}
	sort.Slice(ans, func(a, b int) bool {
		return countMap[ans[a]] > countMap[ans[b]]
	})
	return ans[:k]
}
