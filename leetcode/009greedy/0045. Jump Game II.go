package leetcode

// 给定一个非负整数数组，你最初位于数组的第一个位置。
// 数组中的每个元素代表你在该位置可以跳跃的最大长度。
// 你的目标是使用最少的跳跃次数到达数组的最后一个位置。

// func jump(nums []int) int {
// 	n := len(nums)
// 	if n == 1 {
// 		return 0
// 	}
// 	// 统计两个覆盖范围，当前步的覆盖范围，和下一步的覆盖范围。
// 	// （因为如果和之前的跳跃游戏一样只统计当前步，你就不知道下一步能跳多远）
// 	cur, next := 0, 0
// 	// 步数
// 	step := 0
// 	for i := 0; i < n; i++ {
// 		// 更新下一步的覆盖范围
// 		next = max(nums[i]+i, next)
// 		// 走到当前步的最远覆盖距离了
// 		if i == cur {
// 			// 还没到终点
// 			if cur != n-1 {
// 				step++
// 				cur = next
// 				// 如果下一步能到终点，直接返回步数
// 				if cur >= n-1 {
// 					return step
// 				}

// 			} else {
// 				// 能到终点直接返回步数
// 				return step
// 			}
// 		}
// 	}
// 	return step
// }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 贪心版本二
// 移动下标只要遇到当前覆盖最远距离的下标，直接步数加一，不考虑是不是终点的情况。
func jump(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	cur, next := 0, 0
	step := 0
	// 最大只移动到numsLen-2的地方
	for i := 0; i < n-1; i++ {
		next = max(nums[i]+i, next)
		if i == cur {
			cur = next
			//到达当前覆盖最远距离的下标，直接步数加一，
			// 因为题目保证了一定会到达终点，所以不需要额外的判断了。
			step++
		}
	}
	return step
}
