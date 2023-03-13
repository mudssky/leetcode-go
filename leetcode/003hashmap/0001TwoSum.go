package leetcode

// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。

// 先来最简单的方法，遍历一遍存到map里，第二遍判断是否有和为target的数字
// func twoSum(nums []int, target int) []int {
// 	numMap := make(map[int][]int, 0)
// 	res := []int{}
// 	for i, v := range nums {
// 		numMap[v] = append(numMap[v], i)
// 	}
// 	for i, v := range nums {
// 		want := target - v
// 		if indexList, ok := numMap[want]; ok {
// 			if want == v {

// 				if len(indexList) > 1 {
// 					// 相等时，需要两个不一样的数组
// 					res = []int{i, indexList[0]}
// 					if indexList[0] == i {
// 						res = []int{i, indexList[1]}
// 					}
// 					return res
// 				} else {
// 					continue
// 				}
// 			} else {
// 				res = []int{i, indexList[0]}
// 				return res
// 			}
// 		}

// 	}
// 	return res
// }

// 可以一遍遍历解决
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int, 0)

	for i, v := range nums {
		if index, ok := numMap[target-v]; ok {
			return []int{index, i}
		}
		numMap[v] = i
	}
	return []int{}
}
