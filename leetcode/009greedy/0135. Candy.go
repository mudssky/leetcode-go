package leetcode

// 老师想给孩子们分发糖果，有 N 个孩子站成了一条直线，老师会根据每个孩子的表现，预先给他们评分。

// 你需要按照以下要求，帮助老师给这些孩子分发糖果：
// 每个孩子至少分配到 1 个糖果。
// 相邻的孩子中，评分高的孩子必须获得更多的糖果。
// 那么这样下来，老师至少需要准备多少颗糖果呢？

func candy(ratings []int) int {
	/**先确定一边，再确定另外一边
	    1.先从左到右，当右边的大于左边的就加1
	    2.再从右到左，当左边的大于右边的就再加1
	**/
	// 存放结果需要糖果数的数组
	need := make([]int, len(ratings))
	sum := 0
	// 初始化(每个人至少一个糖果)
	// for i := 0; i < len(ratings); i++ {
	// 	need[i] = 1
	// }
	// 1.先从左到右，当右边的大于左边的就加1
	for i := 0; i < len(ratings)-1; i++ {
		if ratings[i] < ratings[i+1] {
			need[i+1] = need[i] + 1
		}
	}
	// 2.再从右到左，当左边的大于右边的就右边加1，但要花费糖果最少，所以需要做下判断
	for i := len(ratings) - 1; i > 0; i-- {
		if ratings[i-1] > ratings[i] {
			need[i-1] = max(need[i-1], need[i]+1)
		}
	}
	//计算总共糖果
	for i := 0; i < len(ratings); i++ {
		// 前面没有初始化1个糖果，只要在结尾加回来就可以了
		sum += need[i] + 1
	}
	return sum
}
