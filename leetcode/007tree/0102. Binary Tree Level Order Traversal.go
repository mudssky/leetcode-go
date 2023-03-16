package leetcode

// 使用队列
func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil { //防止为空
		return res
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelList := []int{}
		qLen := len(queue)
		for i := 0; i < qLen; i++ {
			front := queue[0]
			queue = queue[1:]
			levelList = append(levelList, front.Val)
			if front.Left != nil {
				queue = append(queue, front.Left)
			}
			if front.Right != nil {
				queue = append(queue, front.Right)
			}
		}
		res = append(res, levelList)
	}
	return res
}

// 递归法
// func levelOrder(root *TreeNode) [][]int {
// 	arr := [][]int{}

// 	depth := 0

// 	var order func(root *TreeNode, depth int)

// 	order = func(root *TreeNode, depth int) {
// 		if root == nil {
// 			return
// 		}
// 		if len(arr) == depth {
// 			arr = append(arr, []int{})
// 		}
// 		arr[depth] = append(arr[depth], root.Val)

// 		order(root.Left, depth+1)
// 		order(root.Right, depth+1)
// 	}

// 	order(root, depth)

// 	return arr
// }
