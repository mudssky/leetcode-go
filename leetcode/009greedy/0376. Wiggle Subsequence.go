package leetcode

// 如果连续数字之间的差严格地在正数和负数之间交替，则数字序列称为摆动序列。
// 第一个差（如果存在的话）可能是正数或负数。少于两个元素的序列也是摆动序列。

// 给定一个整数序列，返回作为摆动序列的最长子序列的长度。
// 通过从原始序列中删除一些（也可以不删除）元素来获得子序列，剩下的元素保持其原始顺序。
func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	ans := 1
	prevDiff := nums[1] - nums[0]
	// 初始两个元素相等的情况下，结果是1
	// 不相等的情况，结果至少是2
	if prevDiff != 0 {
		ans = 2
	}
	for i := 2; i < n; i++ {
		curDiff := nums[i] - nums[i-1]
		// 出现波谷波峰的情况
		// 出现平坡的情况，处理的手法是只记录一侧，即只计算prediff=0的情况。
		// prediff<=0 是平转上升，prediff>=0是平转下降，分别是这两个情况的转折点，需要加到结果里。
		if curDiff > 0 && prevDiff <= 0 || curDiff < 0 && prevDiff >= 0 {
			ans++
			// 出现峰谷时更新prediff，这样就跳过了平坡
			prevDiff = curDiff
		}
	}
	return ans
}
