package leetcode

import "sort"

// 给定一个区间的集合，找到需要移除区间的最小数量，使剩余区间互不重叠。
// 注意: 可以认为区间的终点总是大于它的起点。 区间 [1,2] 和 [2,3] 的边界相互“接触”，但没有相互重叠。

func eraseOverlapIntervals(intervals [][]int) int {
	// 按照右边界升序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	// 保存不重叠的数量
	res := 1
	for i := 1; i < len(intervals); i++ {
		//后一个区间的开始大于前一个区间的结束，说明没有重叠
		if intervals[i][0] >= intervals[i-1][1] {
			res++
		} else {
			// 有重叠的情况（意味着减少一个区间），计算最小重叠区间，用于判断是否和下一个区间重叠，
			// 因为已经确定区间1和区间2重叠，那么区间2或者区间1就是确定移除一个了，下一个区间3，只需要判断区间1和区间2最小的右边界是否重叠。
			intervals[i][1] = min(intervals[i-1][1], intervals[i][1])
		}
	}
	return len(intervals) - res
}
