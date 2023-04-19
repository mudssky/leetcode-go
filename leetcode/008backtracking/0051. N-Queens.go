package leetcode

import "strings"

func solveNQueens(n int) [][]string {
	var res [][]string
	// 二维数组作为棋盘
	chessboard := make([][]string, n)
	for i := 0; i < n; i++ {
		chessboard[i] = make([]string, n)
	}
	// 棋盘初始化，空白为.
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			chessboard[i][j] = "."
		}
	}

	var backtrack func(int)
	backtrack = func(row int) {
		// 遍历完棋盘最后一行，此时回溯应该结束。
		if row == n {
			temp := make([]string, n)
			for i, rowStr := range chessboard {
				temp[i] = strings.Join(rowStr, "")
			}
			res = append(res, temp)
			return
		}
		// 横向遍历
		for i := 0; i < n; i++ {
			if isValid(n, row, i, chessboard) {
				// 经验证合法，可以放入Q
				chessboard[row][i] = "Q"
				// 纵向递归
				backtrack(row + 1)
				// 回溯
				chessboard[row][i] = "."
			}
		}
	}
	backtrack(0)
	return res
}

// 验证Queen是否能放在此处
func isValid(n, row, col int, chessboard [][]string) bool {
	// 同行不需要检验，因为没层遍历的的时候都只有一个
	// 同列不能有Q
	for i := 0; i < row; i++ {
		if chessboard[i][col] == "Q" {
			return false
		}
	}
	// 45度角方向不能有Q
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if chessboard[i][j] == "Q" {
			return false
		}
	}
	// 135度角方向不能有Q
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if chessboard[i][j] == "Q" {
			return false
		}
	}
	return true
}
