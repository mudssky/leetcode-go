package leetcode

import "fmt"

// 给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
// 不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并原地修改输入数组。
// 元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
func removeElement(nums []int, val int) int {
	length := len(nums)
	left := 0
	for i := 0; i < length; i++ {
		if nums[i] != val {
			nums[left] = nums[i]
			left++
		}
	}
	// fmt.Println("nums1", nums)
	return left
}
func main2() {
	test1 := []int{3, 2, 2, 3}
	newLen := removeElement(test1, 2)
	fmt.Println(test1)
	fmt.Println(newLen)
}
