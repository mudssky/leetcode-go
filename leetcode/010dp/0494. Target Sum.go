package leetcode

import "math"

// 给定一个非负整数数组，a1, a2, ..., an, 和一个目标数，S。现在你有两个符号 + 和 -。对于数组中的任意一个整数，你都可以从 + 或 -中选择一个符号添加在前面。

// 返回可以使最终数组和为目标数 S 的所有添加符号的方法数。

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if abs(target) > sum {
		return 0
	}
	// 奇数不能被2整除，此时没有方案。
	if (sum+target)%2 == 1 {
		return 0
	}
	// 计算背包大小
	bag := (sum + target) / 2
	// 定义dp数组
	dp := make([]int, bag+1)
	// 初始化
	dp[0] = 1
	// 遍历顺序
	for i := 0; i < len(nums); i++ {
		for j := bag; j >= nums[i]; j-- {
			//推导公式
			dp[j] += dp[j-nums[i]]
			//fmt.Println(dp)
		}
	}
	return dp[bag]
}

func abs(x int) int {
	return int(math.Abs(float64(x)))
}
