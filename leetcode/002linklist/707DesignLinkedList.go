package leetcode

// type Node struct {
// 	Val  int
// 	Next *Node
// 	Prev *Node
// }
// type MyLinkedList struct {
// 	head   *Node
// 	length int
// }

// func Constructor() MyLinkedList {
// 	return MyLinkedList{&Node{}, 0}
// }

// func (this *MyLinkedList) Get(index int) int {
// 	if this.length < 1 {
// 		return -1
// 	}
// 	currentIndex := 0
// 	cur := this.head
// 	for cur != nil {
// 		if currentIndex == index {
// 			return cur.Val
// 		}
// 		currentIndex++
// 		cur = cur.Next
// 	}
// 	return -1
// }

// func (this *MyLinkedList) AddAtHead(val int) {
// 	// 链表为空的情况特殊处理
// 	if this.length < 1 {
// 		this.head = &Node{val, nil, nil}
// 		this.length = 1
// 		return
// 	}
// 	// 在头部插入
// 	this.head.Prev = &Node{val, this.head, nil}
// 	this.head = this.head.Prev
// 	this.length++
// }

// func (this *MyLinkedList) AddAtTail(val int) {
// 	if this.length < 1 {
// 		this.head = &Node{val, nil, nil}
// 		this.length = 1
// 		return
// 	}
// 	cur := this.head
// 	for cur != nil && cur.Next != nil {
// 		cur = cur.Next
// 	}
// 	cur.Next = &Node{val, nil, cur}
// 	this.length++
// }

// func (this *MyLinkedList) AddAtIndex(index int, val int) {
// 	if index <= 0 {
// 		this.AddAtHead(val)
// 		return
// 	}
// 	if index > this.length {
// 		return
// 	}
// 	if index == this.length {
// 		this.AddAtTail(val)
// 		// newNode := &Node{val, cur, cur.Prev}
// 		// cur.Prev.Next = newNode
// 		// // cur.Next = nil

// 		// this.length++
// 		return
// 	}
// 	currentIndex := 0
// 	cur := this.head
// 	for cur != nil && cur.Next != nil {
// 		if currentIndex == index {
// 			newNode := &Node{val, nil, nil}
// 			newNode.Next = cur
// 			newNode.Prev = cur.Prev
// 			cur.Prev.Next = newNode
// 			cur.Prev = newNode
// 			this.length++
// 			return
// 		}

// 		currentIndex++
// 		cur = cur.Next
// 	}

// 	newNode := &Node{val, cur, cur.Prev}
// 	cur.Prev.Next = newNode
// 	cur.Prev = newNode
// 	this.length++
// }

// func (this *MyLinkedList) DeleteAtIndex(index int) {

// 	cur := this.head

// 	if index == 0 {
// 		this.head = this.head.Next
// 		if this.head != nil {
// 			this.head.Prev = nil
// 		}
// 		this.length--
// 		return
// 	}
// 	currentIndex := 0
// 	for cur != nil && cur.Next != nil {
// 		if currentIndex == index {
// 			cur.Prev.Next = cur.Next
// 			cur.Next.Prev = cur.Prev
// 			this.length--
// 			return
// 		}
// 		cur = cur.Next
// 		currentIndex++
// 	}
// 	if index == this.length-1 {

// 		cur.Prev.Next = nil
// 		cur.Prev = nil
// 		this.length--
// 		return
// 	}

// }
// func (m *MyLinkedList) DisplayList() {
// 	res := []int{}
// 	cur := m.head
// 	for cur != nil {
// 		res = append(res, cur.Val)
// 		cur = cur.Next
// 	}
// 	fmt.Println("res", res, "len", m.length)
// }

// func (m *MyLinkedList) ToList() []int {
// 	res := []int{}
// 	cur := m.head
// 	for cur != nil {
// 		res = append(res, cur.Val)
// 		cur = cur.Next
// 	}
// 	return res
// }

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

// 下面是一个循环双链表的版本
//循环双链表
type MyLinkedList struct {
	dummy *Node
}

type Node struct {
	Val  int
	Next *Node
	Pre  *Node
}

//仅保存哑节点，pre-> rear, next-> head
/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	rear := &Node{
		Val:  -1,
		Next: nil,
		Pre:  nil,
	}
	rear.Next = rear
	rear.Pre = rear
	return MyLinkedList{rear}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	head := this.dummy.Next
	//head == this, 遍历完全
	for head != this.dummy && index > 0 {
		index--
		head = head.Next
	}
	//否则, head == this, 索引无效
	if 0 != index {
		return -1
	}
	return head.Val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	dummy := this.dummy
	node := &Node{
		Val: val,
		//head.Next指向原头节点
		Next: dummy.Next,
		//head.Pre 指向哑节点
		Pre: dummy,
	}

	//更新原头节点
	dummy.Next.Pre = node
	//更新哑节点
	dummy.Next = node
	//以上两步不能反
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	dummy := this.dummy
	rear := &Node{
		Val: val,
		//rear.Next = dummy(哑节点)
		Next: dummy,
		//rear.Pre = ori_rear
		Pre: dummy.Pre,
	}

	//ori_rear.Next = rear
	dummy.Pre.Next = rear
	//update dummy
	dummy.Pre = rear
	//以上两步不能反
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	head := this.dummy.Next
	//head = MyLinkedList[index]
	for head != this.dummy && index > 0 {
		head = head.Next
		index--
	}
	if index > 0 {
		return
	}
	node := &Node{
		Val: val,
		//node.Next = MyLinkedList[index]
		Next: head,
		//node.Pre = MyLinkedList[index-1]
		Pre: head.Pre,
	}
	//MyLinkedList[index-1].Next = node
	head.Pre.Next = node
	//MyLinkedList[index].Pre = node
	head.Pre = node
	//以上两步不能反
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	//链表为空
	if this.dummy.Next == this.dummy {
		return
	}
	head := this.dummy.Next
	//head = MyLinkedList[index]
	for head.Next != this.dummy && index > 0 {
		head = head.Next
		index--
	}
	//验证index有效
	if index == 0 {
		//MyLinkedList[index].Pre = index[index-2]
		head.Next.Pre = head.Pre
		//MyLinedList[index-2].Next = index[index]
		head.Pre.Next = head.Next
		//以上两步顺序无所谓
	}
}
