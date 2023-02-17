package lo

// 返回一个拼接好的新数组
func Concat[T any](collection []T, values ...T) []T {
	result := make([]T, 0, len(collection)+len(values))
	result = append(result, collection...)
	result = append(result, values...)
	return result
}

// 生成第一个数组中独有的元素组成的新数组
// collection是检查的数组,excludes是需要排除的值的数组
func Difference[T comparable](collection []T, excludes []T) []T {
	result := []T{}
	// 初始化需要排除的map,便于查找
	excludesMap := map[T]struct{}{}
	for _, elem := range excludes {
		excludesMap[elem] = struct{}{}
	}

	// 遍历第一个数组,放入第二个数组中不存在的元素,
	for _, elem := range collection {
		if _, ok := excludesMap[elem]; !ok {
			result = append(result, elem)
		}
	}

	return result
}

// 生成第一个数组中独有的元素组成的新数组,对比较一个元素调用iteratee转换后的结果
func DifferenceBy[T comparable, R comparable](collection []T, excludes []T, iteratee func(item T) R) []T {
	result := []T{}
	// 初始化需要排除的map,便于查找
	excludesMap := map[R]struct{}{}
	for _, elem := range excludes {
		exclude := iteratee(elem)
		excludesMap[exclude] = struct{}{}
	}
	// 遍历第一个数组,放入第二个数组中不存在的元素,
	for _, elem := range collection {
		compareElem := iteratee(elem)
		if _, ok := excludesMap[compareElem]; !ok {
			result = append(result, elem)
		}
	}

	return result
}

// 和Difference一样,但是用一个比较函数判断
func DifferenceWith[T comparable, R comparable](collection []T, excludes []R, comparator func(left T, right R) bool) []T {
	result := []T{}
	// 缓存map,把已经比较的结果存起来,加速重复的比较,通过的是true
	cacheMap := map[T]bool{}

	// 遍历第一个数组,放入第二个数组中不存在的元素,
	for _, elem := range collection {
		// 用缓存提升执行效率,如果缓存中比较过了通过,那么直接放入到结果中
		if isPass, ok := cacheMap[elem]; ok {
			// 不通过的情况跳过
			if !isPass {
				continue
			}
			result = append(result, elem)
		}
		isInclude := false
		// 缓存中没有比较过,遍历excludes比较
		for _, ex := range excludes {
			isEqual := comparator(elem, ex)
			if isEqual {
				isInclude = true
				break
			}
		}
		if isInclude {
			cacheMap[elem] = false

		} else {
			// 不包含的情况添加
			result = append(result, elem)
			cacheMap[elem] = true
		}
	}

	return result
}
