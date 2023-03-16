package leetcode

import (
	"container/heap"

	"github.com/mudssky/leetcode-go/constraints"
)

// 使用排序函数实现
// func topKFrequent(nums []int, k int) []int {
// 	countMap := make(map[int]int)
// 	for _, v := range nums {
// 		countMap[v]++
// 	}
// 	ans := []int{}
// 	for k := range countMap {
// 		ans = append(ans, k)
// 	}
// 	sort.Slice(ans, func(a, b int) bool {
// 		return countMap[ans[a]] > countMap[ans[b]]
// 	})
// 	return ans[:k]
// }

// 使用go语言官方提供的接口实现小顶堆
// 用堆的好处是，如果使用快排，那么必须维护一样大小的数组
// 用堆，只用维护k大小的堆。(leetcode上不如快排快，可能是堆里面Pop和Push操作里面数组赋值太耗时。。。)
// 用小顶堆，维护k大小的堆，每次取出的是最小的，因此保留的都是更大的。
type IHeap [][2]int

func (h IHeap) Len() int {
	return len(h)
}

func (h IHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}

func (h IHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IHeap) Push(x any) {
	*h = append(*h, x.([2]int))
}
func (h *IHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	map_num := map[int]int{}
	//记录每个元素出现的次数
	for _, item := range nums {
		map_num[item]++
	}
	h := &IHeap{}
	heap.Init(h)
	//所有元素入堆，堆的长度为k
	// 用小顶堆，每次取出的是最小的，因此保留的都是更大的。
	for key, value := range map_num {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	res := make([]int, k)
	//按顺序返回堆中的元素
	for i := 0; i < k; i++ {
		res[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return res
}

// 堆就是父节点小于等于（小顶堆）子节点的二叉树
// 用数组模拟堆，0是根节点，子节点的下标是父节点*2+1 和+2
// 相反 父节点= (子节点-1)/2 是整除
// down 函数是向下调整某个节点位置的函数，它的作用是将该节点与其左右子节点中较小的一个交换位置，直到满足最小堆性质或者到达叶子节点为止。down
// 函数通常用于删除堆顶元素后，将新的根节点向下调整到合适的位置。
// up 函数是向上调整某个节点位置的函数，它的作用是将该节点与其父节点比较大小，如果该节点比父节点小，则交换位置，直到满足最小堆性质或者到达根节点为止。up
// 函数通常用于插入新元素后，将新加入的元素向上调整到合适的位置。

func down[T constraints.Ordered](arr []T, index, length int) {
	for {
		leftIndex := index>>1 + 1
		// 交换的节点下标，初始为左子节点
		changeIndex := leftIndex
		if leftIndex >= length {
			break
		}
		rightIndex := leftIndex + 1
		// 用最小的交换
		if arr[rightIndex] < arr[leftIndex] {
			changeIndex = rightIndex
		}
		// 父节点小于子节点，满足最小堆，不用交换了，退出循环
		if arr[index] <= arr[changeIndex] {
			break
		}
		arr[changeIndex], arr[index] = arr[index], arr[changeIndex]
		// 改变父节点
		index = changeIndex
	}
}

// 对每个父节点执行down函数就能完成建堆，因为针对父节点，down执行次数少一些
func heapify[T constraints.Ordered](arr []T) {
	arrLen := len(arr)
	// 从最后一个父节点开始down操作
	for i := arrLen/2 - 1; i >= 0; i-- {
		down(arr, i, arrLen)
	}
}
