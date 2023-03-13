package leetcode

// func intersection(nums1 []int, nums2 []int) []int {
// 	numMap := make(map[int]bool, len(nums1))
// 	res := []int{}
// 	for _, v := range nums1 {
// 		numMap[v] = true
// 	}
// 	resMap := make(map[int]bool)
// 	for _, v := range nums2 {
// 		if _, ok := numMap[v]; ok {
// 			if _, ok := resMap[v]; !ok {
// 				res = append(res, v)
// 				resMap[v] = true
// 			}

// 		}
// 	}
// 	return res
// }

// 不需要第二个map，可以匹配到一次后删除，这样避免重复

func intersection(nums1 []int, nums2 []int) []int {
	numMap := make(map[int]bool, len(nums1))
	res := make([]int, 0)
	for _, v := range nums1 {
		numMap[v] = true
	}

	for _, v := range nums2 {
		if _, ok := numMap[v]; ok {
			res = append(res, v)
			delete(numMap, v)
		}
	}
	return res
}
