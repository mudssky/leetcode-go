package leetcode

import "math"

// func getMinimumDifference(root *TreeNode) int {
// 	// 转换为求有序数组两点间差绝对值的最小值

// 	arr := inorderTraversal(root)
// 	res := getAbs(arr[1] - arr[0])
// 	for i := 0; i < len(arr)-1; i++ {
// 		diff := getAbs(arr[i+1] - arr[i])
// 		if diff < res {
// 			res = diff
// 		}
// 	}
// 	return res
// }

// func minAbs[T constraints.Ordered](a,b T){
// 	positiveA,positiveB:=a,b
// 	if a<0{
// 		positiveA=-a
// 	}
// 	if  b<0{
// 		positiveB=-b
// 	}

//  if (a<b){
// 	return a
// }
// return b
// }

func getAbs(a int) int {
	var res int
	if a < 0 {
		res = -a
	} else {
		res = a
	}
	return res
}

// 可以在中序遍历的同时计算最小值
func getMinimumDifference(root *TreeNode) int {
	// 保留前一个节点的指针
	var prev *TreeNode
	// 定义一个比较大的值
	min := math.MaxInt64
	var travel func(node *TreeNode)
	travel = func(node *TreeNode) {
		if node == nil {
			return
		}
		travel(node.Left)
		if prev != nil && node.Val-prev.Val < min {
			min = node.Val - prev.Val
		}
		prev = node
		travel(node.Right)
	}
	travel(root)
	return min
}
