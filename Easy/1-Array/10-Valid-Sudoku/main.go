package main

import "fmt"

func main() {
	board1 := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	board2 := [][]byte{
		{'.', '8', '7', '6', '5', '4', '3', '2', '1'},
		{'2', '.', '8', '.', '.', '.', '.', '.', '.'},
		{'3', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'4', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'5', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'6', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'7', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'8', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'9', '.', '.', '.', '.', '.', '.', '.', '.'},
	}

	out1, out2 := isValidSudoku(board1), isValidSudoku(board2)

	fmt.Println(out1, out2)
}

func isValidSudoku(board [][]byte) bool {
	rows := make([]map[byte]bool, 9)
	columns := make([]map[byte]bool, 9)
	squares := make([]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool, 9)
		columns[i] = make(map[byte]bool, 9)
		squares[i] = make(map[byte]bool, 9)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			num := board[i][j]

			if num == '.' {
				continue
			}

			squareIndex := (i/3)*3 + j/3

			if rows[i][num] {
				fmt.Printf("Duplicate %c found in row %d\n", num, i+1)
				return false
			}
			if columns[j][num] {
				fmt.Printf("Duplicate %c found in column %d\n", num, j+1)
				return false
			}
			if squares[squareIndex][num] {
				fmt.Printf("Duplicate %c found in square %d\n", num, squareIndex+1)
				return false
			}

			rows[i][num] = true
			columns[j][num] = true
			squares[squareIndex][num] = true
		}
	}

	return true
}
