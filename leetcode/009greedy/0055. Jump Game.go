package leetcode

// 给定一个非负整数数组，你最初位于数组的第一个位置。
// 数组中的每个元素代表你在该位置可以跳跃的最大长度。
// 判断你是否能够到达最后一个位置。

func canJump(nums []int) bool {
	numsLen := len(nums)
	// 只有一个元素的情况，能达到
	if numsLen == 1 {
		return true
	}
	// 只要判断最后的跳跃范围能够覆盖到终点就行了。
	cover := 0
	for i := 0; i <= cover; i++ {
		distance := i + nums[i]
		if distance > cover {
			cover = distance
		}
		if cover >= numsLen-1 {
			return true
		}
	}
	return false

}
