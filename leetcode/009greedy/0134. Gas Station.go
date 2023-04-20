package leetcode

// 在一条环路上有 N 个加油站，其中第 i 个加油站有汽油 gas[i] 升。
// 你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。你从其中的一个加油站出发，开始时油箱为空。
// 如果你可以绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1。

// 说明:
// 如果题目有解，该答案即为唯一答案。
// 输入数组均为非空数组，且长度相同。
// 输入数组中的元素均为非负数。

// 暴力解法
func canCompleteCircuit2(gas []int, cost []int) int {
	// gas 是加油站列表，数值是每个加油站可获得的油量
	// cost是从i个加油站到i+1个加油站需要消耗的汽油

	// 遍历每个加油站作为起点的情况
	for i := 0; i < len(gas); i++ {
		// 先开到下一个加油站来初始化变量
		// 剩余油量
		rest := gas[i] - cost[i]
		// i超过cost列表以后，是按照取模计算油耗的。。。题目没说
		// +1取模的结果正好是下一个元素，也就是数组第一个元素
		index := (i + 1) % len(cost)
		for rest > 0 && index != i {
			// 油耗计算
			rest += gas[index] - cost[index]
			// 去下一个加油站
			index = (index + 1) % len(cost)
		}
		if rest >= 0 && index == i {
			return i
		}
	}
	return -1
}

// 贪心算法
// 计算经过每个加油站的剩余油耗，遍历累加到curSum.如果curSum小于0，说明这一段里面油都是不够用的，
// 起点调到后面
func canCompleteCircuit(gas []int, cost []int) int {
	curSum := 0
	// 总剩余油量是负数，说明油是不够一圈的。
	totalSum := 0
	start := 0
	for i := 0; i < len(gas); i++ {
		curSum += gas[i] - cost[i]
		totalSum += gas[i] - cost[i]
		// curSum，累计剩余油量，如果为负数,说明这一段的位置作为起点都开不了一圈。。。。
		if curSum < 0 {
			start = i + 1
			curSum = 0
		}
	}
	if totalSum < 0 {
		return -1
	}
	return start
}
