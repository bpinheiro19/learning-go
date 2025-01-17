package main

import (
	"bytes"
	"log"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Game struct {
	mplusFaceSource *text.GoTextFaceSource
	board           [][]Tile
	running         bool
	gameover        bool
	win             bool
}

const (
	ScreenWidth  = 1000
	ScreenHeight = 850
	boardWidth   = 5
	boardHeight  = 6
)

func newGame() *Game {
	game := &Game{}
	game.init()
	game.initBoard()
	game.printBoard()
	game.running = true
	game.gameover = false
	game.win = false
	return game
}

func (g *Game) init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}
	g.mplusFaceSource = s
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenHeight, ScreenWidth
}

func main() {
	ebiten.SetWindowSize(ScreenHeight, ScreenWidth)
	ebiten.SetWindowTitle("Wordle")
	if err := ebiten.RunGame(newGame()); err != nil {
		log.Fatal(err)
	}
}
