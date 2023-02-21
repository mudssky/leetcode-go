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
	count := 0
	// for startX < endX && startY < endY {
	for i := startX; i < endX; i++ {
		res[0][i] = count
		count++
	}
	for i := startY; i < endY; i++ {
		res[i][endX] = count
		count++
	}

	// for i:=endX;>=0
	// }
	// fmt.Println()
	return res
}
