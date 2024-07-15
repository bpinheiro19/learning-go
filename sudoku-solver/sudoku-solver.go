package main

import (
	"fmt"
)

type SudokuBoard struct {
	board [][]uint8
}

func newSudokuBoard() *SudokuBoard {
	board := &SudokuBoard{}
	board.initBoard()
	return board
}

func (sb *SudokuBoard) initBoard() {
	sb.board = [][]uint8{
		{0, 0, 0, 2, 6, 0, 7, 0, 1},
		{6, 8, 0, 0, 7, 0, 0, 9, 0},
		{1, 9, 0, 0, 0, 4, 5, 0, 0},
		{8, 2, 0, 1, 0, 0, 0, 4, 0},
		{0, 0, 4, 6, 0, 2, 9, 0, 0},
		{0, 5, 0, 0, 0, 3, 0, 2, 8},
		{0, 0, 9, 3, 0, 0, 0, 7, 4},
		{0, 4, 0, 0, 5, 0, 0, 3, 6},
		{7, 0, 3, 0, 1, 8, 0, 0, 0},
	}
}

func (sb *SudokuBoard) printBoard() {
	for i := 0; i < 9; i++ {
		fmt.Println(" ------------------------------------- ")
		for n := 0; n < 9; n++ {
			fmt.Print(" | ")
			fmt.Print(sb.board[i][n])

		}
		fmt.Println(" | ")
	}
}

func main() {
	board := newSudokuBoard()
	board.printBoard()
	print(board, "test")
}
