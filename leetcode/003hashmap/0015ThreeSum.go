package leetcode

import (
	"sort"
	"strconv"
	"strings"
)

// 有超时的风险
func threeSum2(nums []int) [][]int {
	numsLen := len(nums)
	numsMap := make(map[int]int, 0)
	for i, v := range nums {
		numsMap[v] = i
	}
	res := [][]int{}
	filterMap := make(map[string]bool, 0)
	for i := 0; i < numsLen; i++ {
		for j := i + 1; j < numsLen; j++ {
			ij := nums[i] + nums[j]
			if index, ok := numsMap[-ij]; ok && j != index && index != i {
				newList := []int{nums[i], nums[j], nums[index]}
				sort.Ints(newList)
				indexStrList := []string{strconv.Itoa(newList[0]), strconv.Itoa(newList[1]), strconv.Itoa(newList[2])}
				// for _, v := range newList {
				// 	indexStrList = append(indexStrList, strconv.Itoa(v))
				// }
				filterString := strings.Join(indexStrList, ",")
				if _, ok := filterMap[filterString]; !ok {
					res = append(res, newList)
					filterMap[filterString] = true
				}

			}
		}
	}
	return res
}

func threeSum(nums []int) [][]int {
	numsLen := len(nums)
	sort.Ints(nums)
	res := [][]int{}
	for i := 0; i < numsLen-2; i++ {
		n1 := nums[i]

		// 排序之后如果第一个元素已经大于零，那么无论如何组合都不可能凑成三元组，直接返回结果就可以了
		if n1 > 0 {
			break
		}
		// 对第一个元素去重
		if i > 0 && n1 == nums[i-1] {
			continue
		}

		l, r := i+1, numsLen-1
		for l < r {
			n2, n3 := nums[l], nums[r]
			sumOf3 := n1 + n2 + n3
			if sumOf3 == 0 {
				res = append(res, []int{n1, n2, n3})
				// 去重逻辑应该放在找到一个三元组之后，对b 和 c去重
				for l < r && nums[l] == n2 {
					l++
				}
				for l < r && nums[r] == n3 {
					r--
				}
			} else if sumOf3 < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return res
}
