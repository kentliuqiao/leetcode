package nqueens

import "strings"

func solveNQueens(n int) [][]string {
	var res = [][]string{}
	board := make([][]string, n)
	for i := range board {
		board[i] = make([]string, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			board[i][j] = "."
		}
	}
	var backTrack func(int)
	backTrack = func(row int) {
		if row == n {
			temp := []string{}
			for _, v := range board {
				temp = append(temp, strings.Join(v, ""))
			}
			res = append(res, temp)
		}
		for col := 0; col < n; col++ {
			if !isValid(board, row, col) {
				continue
			}
			// make a choice
			board[row][col] = "Q"
			// enter next level in the decision tree
			backTrack(row + 1)
			// drawback the choice
			board[row][col] = "."
		}
	}
	backTrack(0)
	return res
}

func isValid(board [][]string, row, col int) bool {
	// check if another Queen exist in the same column
	for i := 0; i < row; i++ {
		if board[i][col] == "Q" {
			return false
		}
	}
	// check if another Queen exist in the upper left
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == "Q" {
			return false
		}
	}
	//check if another Queen exist in the upper right
	for i, j := row-1, col+1; i >= 0 && j < len(board); i, j = i-1, j+1 {
		if board[i][j] == "Q" {
			return false
		}
	}
	return true
}
