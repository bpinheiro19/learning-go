package main

import (
	"fmt"
)

func (g *Game) initBoard() {
	g.board = make([][]Tile, boardHeight)
	for i := 0; i < boardHeight; i++ {
		tls := make([]Tile, boardWidth)
		g.board[i] = tls
	}
}

func (g *Game) printBoard() {
	fmt.Println(" ### 2048 Board ### ")
	fmt.Println(" ----------------- ")
	for i := 0; i < boardHeight; i++ {
		for n := 0; n < boardWidth; n++ {
			fmt.Print(" | ")
			fmt.Print(g.board[i][n].letter.value)
		}
		fmt.Println(" | ")
	}
	fmt.Println(" ----------------- ")
}
