package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 10 {
		fmt.Println("Error: Args are not 9")
		return
	}

	var board [9][9]int

	for i := 0; i < 9; i++ {
		row := os.Args[i+1]
		if len(row) != 9 {
			fmt.Println("Error: row are not 9")
			return
		}
		for j, c := range row {
			if c == '.' {
				board[i][j] = 0
			} else if c >= '1' && c <= '9' {
				board[i][j] = int(c - '0')
			} else {
				fmt.Println("Error: Not \".\" Or Number")
				return
			}
		}
	}
	if !isInitialValid(board) {
		fmt.Println("Error: Duplicate values in row, column, or box")
		return
	}
	if solve(&board) {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				fmt.Print(board[i][j])
				if j != 8 {
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}
	} else {
		fmt.Println("Not Solvable")
	}
}

func solve(board *[9][9]int) bool {
	row, col, found := findEmpty(board)
	if !found {
		return true // solved
	}

	for num := 1; num <= 9; num++ {
		if isValid(board, row, col, num) {
			board[row][col] = num
			if solve(board) {
				return true
			}
			board[row][col] = 0 // backtrack
		}
	}
	return false
}

func findEmpty(board *[9][9]int) (int, int, bool) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				return i, j, true
			}
		}
	}
	return -1, -1, false
}

func isValid(board *[9][9]int, row, col, num int) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == num || board[i][col] == num {
			return false
		}
	}

	boxRow, boxCol := (row/3)*3, (col/3)*3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[boxRow+i][boxCol+j] == num {
				return false
			}
		}
	}
	return true
}

func isInitialValid(board [9][9]int) bool {
	for i := 0; i < 9; i++ {
		inRow := make(map[int]bool)
		inCol := make(map[int]bool)
		inBox := make(map[int]bool)

		for j := 0; j < 9; j++ {
			if board[i][j] != 0 {
				if inRow[board[i][j]] {
					return false
				}
				inRow[board[i][j]] = true
			}

			if board[j][i] != 0 {
				if inCol[board[j][i]] {
					return false
				}
				inCol[board[j][i]] = true
			}

			boxRow := 3*(i/3) + j/3
			boxCol := 3*(i%3) + j%3
			if board[boxRow][boxCol] != 0 {
				if inBox[board[boxRow][boxCol]] {
					return false
				}
				inBox[board[boxRow][boxCol]] = true
			}
		}
	}
	return true
}
