package structures

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 方便测试数据生成
var NULL = -1 << 63

// Ints2TreeNode 利用 []int 生成 *TreeNode
func Ints2TreeNode(ints []int) *TreeNode {
	n := len(ints)
	if n == 0 {
		return nil
	}

	root := &TreeNode{
		Val: ints[0],
	}
	//  队列,存放二叉树每一层的节点
	queue := make([]*TreeNode, 1, n*2)
	queue[0] = root

	// 遍历下标
	i := 1
	for i < n {
		// 每次取出队列的第一个节点,然后移除
		node := queue[0]
		queue = queue[1:]

		if i < n && ints[i] != NULL {
			node.Left = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < n && ints[i] != NULL {
			node.Right = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

// Tree2ints 把 *TreeNode 按照行还原成 []int
func Tree2ints(tn *TreeNode) []int {
	res := make([]int, 0, 1024)

	queue := []*TreeNode{tn}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			nd := queue[i]
			if nd == nil {
				res = append(res, NULL)
			} else {
				res = append(res, nd.Val)
				queue = append(queue, nd.Left, nd.Right)
			}
		}
		queue = queue[size:]
	}
	i := len(res)
	for i > 0 && res[i-1] == NULL {
		i--
	}
	return res[:i]
}

// 从二叉树中找到对应值的节点，
// 二叉树最好每个值是不同的，因为只返回找到的第一个值
func (t *TreeNode) FindNode(val int) *TreeNode {
	q := []*TreeNode{}
	q = append(q, t)
	for len(q) > 0 {
		head := q[0]
		q = q[1:]
		if head.Val == val {
			return head
		}
		if head.Left != nil {
			q = append(q, head.Left)
		}
		if head.Right != nil {
			q = append(q, head.Right)
		}

	}
	return nil
}
