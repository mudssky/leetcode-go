package leetcode

import "fmt"

type Node struct {
	Val  int
	Next *Node
	Prev *Node
}
type MyLinkedList struct {
	head   *Node
	length int
}

func Constructor() MyLinkedList {
	return MyLinkedList{&Node{}, 0}
}

func (this *MyLinkedList) Get(index int) int {
	if this.length < 1 {
		return -1
	}
	currentIndex := 0
	cur := this.head
	for cur != nil {
		if currentIndex == index {
			return cur.Val
		}
		currentIndex++
		cur = cur.Next
	}
	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {
	// 链表为空的情况特殊处理
	if this.length < 1 {
		this.head = &Node{val, nil, nil}
		this.length = 1
		return
	}
	// 在头部插入
	this.head.Prev = &Node{val, this.head, nil}
	this.head = this.head.Prev
	this.length++
}

func (this *MyLinkedList) AddAtTail(val int) {
	if this.length < 1 {
		this.head = &Node{val, nil, nil}
		this.length = 1
		return
	}
	cur := this.head
	for cur != nil && cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = &Node{val, nil, cur}
	this.length++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index <= 0 {
		this.AddAtHead(val)
		return
	}
	if index > this.length {
		return
	}
	if index == this.length {
		this.AddAtTail(val)
		// newNode := &Node{val, cur, cur.Prev}
		// cur.Prev.Next = newNode
		// // cur.Next = nil

		// this.length++
		return
	}
	currentIndex := 0
	cur := this.head
	for cur != nil && cur.Next != nil {
		if currentIndex == index {
			newNode := &Node{val, cur, cur.Prev}
			cur.Prev = newNode
			cur.Next = newNode.Next
			this.length++
			return
		}

		currentIndex++
		cur = cur.Next
	}

}

func (this *MyLinkedList) DeleteAtIndex(index int) {

	cur := this.head

	if index == 0 {
		this.head = this.head.Next
		if this.head != nil {
			this.head.Prev = nil
		}
		this.length--
		return
	}
	currentIndex := 0
	for cur != nil && cur.Next != nil {
		if currentIndex == index {
			cur.Prev.Next = cur.Next
			cur.Next.Prev = cur.Prev
			this.length--
			return
		}
		cur = cur.Next
		currentIndex++
	}
	if index == this.length-1 {

		cur.Prev.Next = nil
		cur.Prev = nil
		this.length--
		return
	}

}
func (m *MyLinkedList) DisplayList() {
	res := []int{}
	cur := m.head
	for cur != nil {
		res = append(res, cur.Val)
		cur = cur.Next
	}
	fmt.Println("res", res, "len", m.length)
}

func (m *MyLinkedList) ToList() []int {
	res := []int{}
	cur := m.head
	for cur != nil {
		res = append(res, cur.Val)
		cur = cur.Next
	}
	return res
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
