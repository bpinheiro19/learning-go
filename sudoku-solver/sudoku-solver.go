package main

import (
	"fmt"
	"log"
)

type SudokuBoard struct {
	board [][]int
}

func newSudokuBoard() *SudokuBoard {
	board := &SudokuBoard{}
	board.initBoard()
	board.printBoard()
	return board
}

func (sb *SudokuBoard) initBoard() {
	sb.board = [][]int{
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
	fmt.Println(" ------------------------------------- ")
}

func (sb *SudokuBoard) checkForZeros() bool {
	for i := 0; i < 9; i++ {
		for n := 0; n < 9; n++ {
			if sb.board[i][n] == 0 {
				return true
			}
		}
	}
	return false
}

func (sb *SudokuBoard) returnFirstZero() (int, int) {
	for i := 0; i < 9; i++ {
		for n := 0; n < 9; n++ {
			if sb.board[i][n] == 0 {
				return i, n
			}
		}
	}
	return 0, 0
}

func (sb *SudokuBoard) checkPossibleNumber(number int, x int, y int) bool {

	for i := 0; i < 9; i++ {
		if sb.board[x][i] == number {
			return false
		}

		if sb.board[i][y] == number {
			return false
		}
	}

	xx := x - x%3
	yy := y - y%3

	for i := 0; i < 3; i++ {
		for n := 0; n < 3; n++ {
			if sb.board[xx+i][yy+n] == number {
				return false
			}
		}
	}
	return true
}

func (sb *SudokuBoard) findSolution() {
	if sb.checkForZeros() {
		x, y := sb.returnFirstZero()

		for i := 1; i < 10; i++ {
			if sb.checkPossibleNumber(i, x, y) {
				sb.board[x][y] = i
				sb.findSolution()
				sb.board[x][y] = 0
			}
		}
		return
	}

	if sb.isSolved() {
		log.Print("Sudoku puzzle is solved!!")
		sb.printBoard()
	}
}

func (sb *SudokuBoard) isSolved() bool {

	for i := 0; i < 9; i++ {
		sumx := 0
		sumy := 0

		for n := 0; n < 9; n++ {
			sumx += sb.board[n][i]
		}

		for n := 0; n < 9; n++ {
			sumy += sb.board[i][n]
		}

		if sumx != 45 || sumy != 45 {
			return false
		}
	}

	for a := 0; a < 9; a += 3 {
		sum1 := 0
		sum2 := 0
		sum3 := 0

		for b := 0; b < 3; b++ {
			sum1 = sum1 + sb.board[a+b][0] + sb.board[a+b][1] + sb.board[a+b][2]
			sum2 = sum2 + sb.board[a+b][3] + sb.board[a+b][4] + sb.board[a+b][5]
			sum3 = sum3 + sb.board[a+b][6] + sb.board[a+b][7] + sb.board[a+b][8]
		}

		if sum1 != 45 || sum2 != 45 || sum3 != 45 {
			return false
		}
	}

	return true
}

func main() {
	board := newSudokuBoard()
	board.findSolution()
}
