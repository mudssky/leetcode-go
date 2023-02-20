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

// Intersect returns the intersection between two collections.
func Intersect[T comparable](list1 []T, list2 []T) []T {
	result := []T{}
	seen := map[T]struct{}{}

	for _, elem := range list1 {
		seen[elem] = struct{}{}
	}

	for _, elem := range list2 {
		if _, ok := seen[elem]; ok {
			result = append(result, elem)
		}
	}

	return result
}

// Union returns all distinct elements from given collections.
// result returns will not change the order of elements relatively.
func Union[T comparable](lists ...[]T) []T {
	result := []T{}
	seen := map[T]struct{}{}

	for _, list := range lists {
		for _, e := range list {
			if _, ok := seen[e]; !ok {
				seen[e] = struct{}{}
				result = append(result, e)
			}
		}
	}

	return result
}

// Uniq returns a duplicate-free version of an array, in which only the first occurrence of each element is kept.
// The order of result values is determined by the order they occur in the array.
// Play: https://go.dev/play/p/DTzbeXZ6iEN
func Uniq[T comparable](collection []T) []T {
	result := make([]T, 0, len(collection))
	seen := make(map[T]struct{}, len(collection))

	for _, item := range collection {
		if _, ok := seen[item]; ok {
			continue
		}

		seen[item] = struct{}{}
		result = append(result, item)
	}

	return result
}

// UniqBy returns a duplicate-free version of an array, in which only the first occurrence of each element is kept.
// The order of result values is determined by the order they occur in the array. It accepts `iteratee` which is
// invoked for each element in array to generate the criterion by which uniqueness is computed.
// Play: https://go.dev/play/p/g42Z3QSb53u
func UniqBy[T any, U comparable](collection []T, iteratee func(item T) U) []T {
	result := make([]T, 0, len(collection))
	seen := make(map[U]struct{}, len(collection))

	for _, item := range collection {
		key := iteratee(item)

		if _, ok := seen[key]; ok {
			continue
		}

		seen[key] = struct{}{}
		result = append(result, item)
	}

	return result
}

// Contains returns true if an element is present in a collection.
func Contains[T comparable](collection []T, element T) bool {
	for _, item := range collection {
		if item == element {
			return true
		}
	}

	return false
}

// contains的别名
func Include[T comparable](collection []T, element T) bool {
	return Contains(collection, element)
}

// ContainsBy returns true if predicate function return true.
func ContainsBy[T any](collection []T, predicate func(item T) bool) bool {
	for _, item := range collection {
		if predicate(item) {
			return true
		}
	}

	return false
}

// Every returns true if all elements of a subset are contained into a collection or if the subset is empty.
func Every[T comparable](collection []T, subset []T) bool {
	for _, elem := range subset {
		if !Contains(collection, elem) {
			return false
		}
	}

	return true
}

// EveryBy returns true if the predicate returns true for all of the elements in the collection or if the collection is empty.
func EveryBy[T any](collection []T, predicate func(item T) bool) bool {
	for _, v := range collection {
		if !predicate(v) {
			return false
		}
	}

	return true
}

// Some returns true if at least 1 element of a subset is contained into a collection.
// If the subset is empty Some returns false.
func Some[T comparable](collection []T, subset []T) bool {
	for _, elem := range subset {
		if Contains(collection, elem) {
			return true
		}
	}

	return false
}

// SomeBy returns true if the predicate returns true for any of the elements in the collection.
// If the collection is empty SomeBy returns false.
func SomeBy[T any](collection []T, predicate func(item T) bool) bool {
	for _, v := range collection {
		if predicate(v) {
			return true
		}
	}

	return false
}
