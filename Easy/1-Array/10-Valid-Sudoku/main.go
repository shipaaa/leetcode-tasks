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
	// board2 := [][]byte{
	// 	{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
	// 	{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
	// 	{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
	// 	{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
	// 	{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
	// 	{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
	// 	{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
	// 	{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
	// 	{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	// }

	out1 := isValidSudoku(board1)
	// , isValidSudoku(board2)

	fmt.Println(out1)
}

func isValidSudoku(board [][]byte) bool {
	rows := make([]map[string]bool, 9)
	columns := make([]map[string]bool, 9)
	squares := make([]map[string]bool, 9)

	for i := range 9 {
		rows[i] = make(map[string]bool, 9)
		columns[i] = make(map[string]bool, 9)
		squares[i] = make(map[string]bool, 9)
	}

	for i := 0; i < 1; i++ {
		for j := 0; j < 9; j++ {
			num := string(board[i][j])

			if num == "." {
				continue
			}

			squareIndex := (i/3)*3 + j*3

			if rows[i][num] || columns[i][num] || squares[squareIndex][num] {

				return false
			}

			rows[i][num] = true
			columns[j][num] = true
			squares[squareIndex][num] = true
		}
	}

	return true
}
