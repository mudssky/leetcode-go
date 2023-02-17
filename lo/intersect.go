package lo

// Difference returns the difference between two collections.
// The first value is the collection of element absent of list2.
// The second value is the collection of element absent of list1.
// 两个切片对比,结果返回两个,第一个是list1独有的元素,第二个是list2独有的元素
// 这个有点蠢了,和lodash只遍历第一个数组不同,可能产生多余操作
// func Difference[T comparable](list1 []T, list2 []T) ([]T, []T) {
// 	left := []T{}
// 	right := []T{}

// 	// 初始化两个map
// 	seenLeft := map[T]struct{}{}
// 	seenRight := map[T]struct{}{}

// 	for _, elem := range list1 {
// 		seenLeft[elem] = struct{}{}
// 	}
// 	for _, elem := range list2 {
// 		seenRight[elem] = struct{}{}
// 	}

// 	// 遍历第一个数组,放入第二个数组中不存在的元素,
// 	for _, elem := range list1 {
// 		if _, ok := seenRight[elem]; !ok {
// 			left = append(left, elem)
// 		}
// 	}
// 	// 遍历第二个数组,放入第一个数组中不存在的元素
// 	for _, elem := range list2 {
// 		if _, ok := seenLeft[elem]; !ok {
// 			right = append(right, elem)
// 		}
// 	}

// 	return left, right
// }
