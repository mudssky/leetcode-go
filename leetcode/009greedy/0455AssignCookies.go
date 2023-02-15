package leetcode

import (
	"sort"
)

// 假设你是一位很棒的家长，想要给你的孩子们一些小饼干。但是，每个孩子最多只能给一块饼干。

// 对每个孩子 i，都有一个胃口值 g[i]，这是能让孩子们满足胃口的饼干的最小尺寸；并且每块饼干 j，都有一个尺寸 s[j] 。如果 s[j] >= g[i]，我们可以将这个饼干 j 分配给孩子 i ，这个孩子会得到满足。你的目标是尽可能满足越多数量的孩子，并输出这个最大数值。

// 1. 为了不影响原数组,先拷贝，然后再排序
// 因为排序之后容易进行比较
// 但是实际上leetcode是不检查是否改变原数组的，所以可以在原数组上排序，这样执行用时28ms，击败64%用户
// 时间复杂度相当于快速排序的复杂度O(logn)
// 从小饼干开始喂的算法，遍历了整个饼干数组
// func findContentChildren(g []int, s []int) int {
// 	res := 0
// 	gCopy := util.Copy(g)
// 	sCopy := util.Copy(s)
// 	sort.Ints(gCopy)
// 	sort.Ints(sCopy)
// 	for j := 0; j < len(s) && res < len(g); j++ {
// 		if sCopy[j] >= gCopy[res] {
// 			res += 1
// 		}

//		}
//		return res
//	}
// func findContentChildren(g []int, s []int) int {
// 	sort.Ints(g)
// 	sort.Ints(s)
// 	count := 0
// 	gLen := len(g)
// 	for j := 0; j < len(s); j++ {
// 		if count < gLen && s[j] >= g[count] {
// 			count++
// 		}

// 	}
// 	return count
// }

// 2.其实从大饼干开始喂更快
// 一开始从小饼干开始想，属于是节约的惯性思维了，因为一直穷，所以思想局限了。
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	gLen := len(g)
	// sLen := len(s)
	start, count := len(s)-1, 0
	for i := gLen - 1; i >= 0 && start >= 0; i-- {
		if s[start] >= g[i] {
			count++
			start--
		}
	}
	return count
}
