package main

import (
	"bytes"
	"log"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Game struct {
	mplusFaceSource  *text.GoTextFaceSource
	board            [][]Tile
	availableLetters []Letter
	currentGuess     []Letter
	word             string
	attemptsMade     int
	running          bool
	gameover         bool
	win              bool
}

const (
	ScreenWidth  = 1000
	ScreenHeight = 850
	boardWidth   = 5
	boardHeight  = 6
)

// TODO Implement method to get a random word each day
func getWordleWord() string {
	return "plead"
}

func newGame() *Game {
	game := &Game{}
	game.init()
	game.initBoard()
	game.printBoard()
	game.running = true
	game.gameover = false
	game.win = false
	game.attemptsMade = 0
	game.word = getWordleWord()
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
	if g.running {

		if g.gameover || g.win {
			g.running = false
		}

	} else {

		if repeatingKeyPressed(ebiten.KeyEnter) {
			g.restartGame()
		}
	}

	return nil
}

func (g *Game) restartGame() {
	g.initBoard()
	g.running = true
	g.gameover = false
	g.win = false
}

func repeatingKeyPressed(key ebiten.Key) bool {
	d := inpututil.KeyPressDuration(key)
	return d >= 1 && d < 2
}

func (g *Game) checkGameOver() bool {
	return g.attemptsMade == 6
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
