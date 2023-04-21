package leetcode

import "sort"

// 在二维空间中有许多球形的气球。对于每个气球，提供的输入是水平方向上，气球直径的开始和结束坐标。
// 由于它是水平的，所以纵坐标并不重要，因此只要知道开始和结束的横坐标就足够了。开始坐标总是小于结束坐标。

// 一支弓箭可以沿着 x 轴从不同点完全垂直地射出。在坐标 x 处射出一支箭，若有一个气球的直径的开始和结束坐标为 xstart，xend， 且满足  xstart ≤ x ≤ xend，则该气球会被引爆。
// 可以射出的弓箭的数量没有限制。 弓箭一旦被射出之后，可以无限地前进。我们想找到使得所有气球全部被引爆，所需的弓箭的最小数量。

// 给你一个数组 points ，其中 points [i] = [xstart,xend] ，返回引爆所有气球所必须射出的最小弓箭数。

func findMinArrowShots(points [][]int) int {
	var res int = 1 //弓箭数
	//先按照第一位排序
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	for i := 1; i < len(points); i++ {
		if points[i-1][1] < points[i][0] { //如果前一位的右边界小于后一位的左边界，则一定不重合
			res++
		} else {
			points[i][1] = min(points[i-1][1], points[i][1]) // 更新重叠气球最小右边界,覆盖该位置的值，留到下一步使用
		}
	}
	return res
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
