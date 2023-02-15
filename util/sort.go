package util

// 最基础的冒泡排序
func BubbleSort[T Number](arr []T) {
	length := len(arr)
	for i := 0; i < length; i++ {
		// 每次循环末尾都固定一个最大值，所以下次遍历的数组长度减1
		for j := 0; j < length-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// 最基础的插入排序
func InsertSort[T Number](arr []T) {
	length := len(arr)
	for i := 1; i < length; i++ {

		element := arr[i]
		j := i - 1
		//之前的元素如果比element大，那么将数组使用这种交换的方法整体后移，直到比element小的位置就是插入点
		for j >= 0 && arr[j] > element {
			arr[j+1] = arr[j]
			j--
		}
		// 如果之前的元素不比element大，当前就是合适的位置
		arr[j+1] = element
	}
}
