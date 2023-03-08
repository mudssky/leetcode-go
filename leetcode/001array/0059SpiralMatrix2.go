package leetcode

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	endX := n - 1
	endY := endX
	startX := 0
	startY := 0
	count := 1
	loop := n / 2
	for loop > 0 {
		for i := startX; i < endX; i++ {
			res[startX][i] = count
			count++
		}
		// fmt.Println("1 ", res)
		for i := startY; i < endY; i++ {
			res[i][endX] = count
			count++
		}
		// fmt.Println("2 ", res)
		for i := endX; i > startX; i-- {
			res[endY][i] = count
			count++
		}
		// fmt.Println("3 ", res)
		for i := endY; i > startY; i-- {
			res[i][startX] = count
			count++
		}
		startX++
		startY++
		endX--
		endY--
		loop--
	}
	mid := n / 2
	if n%2 != 0 {
		res[mid][mid] = n * n
	}
	// for i:=endX;>=0
	// }
	// fmt.Println("4 ", res)
	return res
}
