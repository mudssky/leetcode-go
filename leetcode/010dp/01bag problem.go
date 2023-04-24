package leetcode

import "log"

// 有n件物品和一个最多能背重量为w 的背包。第i件物品的重量是weight[i]，得到的价值是value[i] 。每件物品只能用一次，求解将哪些物品装入背包里物品价值总和最大。
func zeroOneBag(weight, value []int, bagweight int) int {
	// dp 是背包负重和物品重量的二维数组，
	// dp[i][j]的含义：从下标为[0-i]的物品里任意取，放进容量为j的背包，价值总和最大是多少。
	dp := make([][]int, len(weight))
	for i := range dp {
		dp[i] = make([]int, bagweight+1)
	}
	// 初始化dp
	// 背包负重为0的情况，价值一定为0
	// 因为数组初始值就是0，所以不用手动赋值了。

	// 背包负重比编号为0第一个物品大的情况，此时i=0的情况，一定是编号0物品的价值
	for j := bagweight; j >= weight[0]; j-- {
		dp[0][j] = value[0]
	}
	// 其他状态可以由状态转移方程推出
	for i := 1; i < len(weight); i++ {
		for j := 0; j <= bagweight; j++ {
			// 背包放不下的情况下，和不放价值相同
			if j < weight[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				// 能放入的情况下，需要和不放的情况比较，选取最大的。
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
			}
		}
	}
	return dp[len(weight)-1][bagweight]

}

// 压缩到1维数组
func zeroOneBag2(weight, value []int, bagweight int) int {
	// dp[j] 表示背包负重为j的情况下，能容纳的最大物品价值
	dp := make([]int, bagweight+1)
	for i := 0; i < len(weight); i++ {
		// 这里必须倒序,区别二维,因为二维dp保存了i的状态
		// 倒序遍历是为了保证物品i只被放入一次
		for j := bagweight; j >= weight[i]; j-- {
			dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
		}
	}
	return dp[bagweight]

}

// 有N件物品和一个最多能背重量为W的背包。第i件物品的重量是weight[i]，得到的价值是value[i] 。
// 每件物品都有无限个（也就是可以放入背包多次），求解将哪些物品装入背包里物品价值总和最大。
// 完全背包和01背包问题唯一不同的地方就是，每种物品有无限件。

// 先遍历物品，再遍历背包负重
func complete01Bag(weight, value []int, bagWeight int) int {
	// 定义dp数组 和初始化
	dp := make([]int, bagWeight+1)
	// 遍历顺序
	for i := 0; i < len(weight); i++ {
		// 正序会多次添加 value[i]
		for j := weight[i]; j <= bagWeight; j++ {
			// 推导公式
			dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
			// debug
			//fmt.Println(dp)
		}
	}
	return dp[bagWeight]
}

// 先遍历背包负重
func complete01Bag02(weight, value []int, bagWeight int) int {
	// 定义dp数组 和初始化
	dp := make([]int, bagWeight+1)
	// 遍历顺序
	// j从0 开始
	for j := 0; j <= bagWeight; j++ {
		for i := 0; i < len(weight); i++ {
			if j >= weight[i] {
				// 推导公式
				dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
			}
			// debug
			//fmt.Println(dp)
		}
	}
	return dp[bagWeight]
}

// 有N种物品和一个容量为V 的背包。第i种物品最多有Mi件可用，每件耗费的空间是Ci ，价值是Wi 。求解将哪些物品装入背包可使这些物品的耗费的空间 总和不超过背包容量，且价值总和最大。
// 在01背包基础上增加了物品的数量
// 多重背包可以化解为 01 背包
func multiplePack(weight, value, nums []int, bagWeight int) int {
	// 转化为01背包问题
	for i := 0; i < len(nums); i++ {
		for nums[i] > 1 {
			weight = append(weight, weight[i])
			value = append(value, value[i])
			nums[i]--
		}
	}
	log.Println(weight)
	log.Println(value)

	res := make([]int, bagWeight+1)
	for i := 0; i < len(weight); i++ {
		for j := bagWeight; j >= weight[i]; j-- {
			res[j] = max(res[j], res[j-weight[i]]+value[i])
		}

	}

	return res[bagWeight]
}
