package main

import (
	"bytes"
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"strconv"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Game struct {
	tiles [][]Tile
	keys  []ebiten.Key
}

const (
	ScreenWidth  = 700
	ScreenHeight = 700
	boardSize    = 4

	tileSize    = 150
	tileMargin  = 160
	tileSpacing = 50

	fontSize = 90
)

var (
	mplusFaceSource *text.GoTextFaceSource
)

func newGame() *Game {
	game := &Game{}
	game.init()
	game.initBoard()
	game.printBoard()
	return game
}

func (g *Game) init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}
	mplusFaceSource = s
}

func (g *Game) initBoard() {
	g.tiles = make([][]Tile, 4)

	for i := 0; i < boardSize; i++ {
		tls := make([]Tile, 4)
		for n := 0; n < boardSize; n++ {
			tls[n] = Tile{value: 0, x: i, y: n}
		}
		g.tiles[i] = tls
	}
}

func (g *Game) printBoard() {
	fmt.Println(" ### 2048 Board ### ")
	fmt.Println(" ----------------- ")
	for i := 0; i < boardSize; i++ {
		for n := 0; n < boardSize; n++ {
			fmt.Print(" | ")
			fmt.Print(g.tiles[i][n].value)

		}
		fmt.Println(" | ")
	}
	fmt.Println(" ----------------- ")
}

func (g *Game) Update() error {

	if repeatingKeyPressed(ebiten.KeyUp) {
		g.spawnTile()
	}

	if repeatingKeyPressed(ebiten.KeyDown) {
		g.spawnTile()
	}

	if repeatingKeyPressed(ebiten.KeyLeft) {
		g.spawnTile()
	}

	if repeatingKeyPressed(ebiten.KeyRight) {
		g.spawnTile()
	}

	//g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	/*
		for _, p := range g.keys {
			if p == ebiten.KeyUp {
				g.spawnTile()
			}
			fmt.Println(p)

		}*/

	return nil
}

func repeatingKeyPressed(key ebiten.Key) bool {
	d := inpututil.KeyPressDuration(key)
	return d >= 1 && d < 2
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(BACKGROUND_COLOR)

	for i := 0; i < boardSize+1; i++ {
		for n := 0; n < boardSize; n++ {
			line := ebiten.NewImage(3, 153)
			line.Fill(color.Black)
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(i*tileSize+tileSpacing), float64(n*tileSize+tileSpacing))
			screen.DrawImage(line, op)

		}
	}

	for i := 0; i < boardSize; i++ {
		for n := 0; n < boardSize+1; n++ {
			line := ebiten.NewImage(150, 3)
			line.Fill(color.Black)
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(i*tileSize+tileSpacing), float64(n*tileSize+tileSpacing))
			screen.DrawImage(line, op)
		}
	}

	for i := 0; i < boardSize; i++ {
		for n := 0; n < boardSize; n++ {
			if g.tiles[i][n].value != 0 {
				tile := ebiten.NewImage(tileSize-4, tileSize-4)
				tile.Fill(COLOR2)
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(i*tileSize+tileSpacing+3), float64(n*tileSize+tileSpacing+3))
				screen.DrawImage(tile, op)

				optext := &text.DrawOptions{}
				optext.GeoM.Translate(float64(i*tileSize+tileSpacing+45), float64(n*tileSize+tileSpacing+15))
				optext.ColorScale.ScaleWithColor(color.Black)
				text.Draw(screen, strconv.Itoa(g.tiles[i][n].value), &text.GoTextFace{
					Source: mplusFaceSource,
					Size:   float64(fontSize),
				}, optext)
			}
		}
	}

}

func (g *Game) moveTile(x, y, w, z int) {
	val := g.tiles[x][y].value
	g.tiles[x][y].value = 0
	g.tiles[w][z].value = val
}

func (g *Game) spawnTile() {
	valid := false

	for !valid {

		x := rand.Intn(4)
		y := rand.Intn(4)

		val := func() int {
			if rand.Intn(10) < 9 {
				return 2
			}
			return 4
		}()

		if g.tiles[x][y].value == 0 {
			valid = true
			g.tiles[x][y].value = val
		}
	}

}

func (g *Game) checkWin() bool {
	for i := 0; i < boardSize; i++ {
		for n := 0; n < boardSize; n++ {
			if g.tiles[i][n].hasMaxValue() {
				return true
			}
		}
	}
	return false
}

func (g *Game) checkGameOver() bool {
	for i := 0; i < boardSize; i++ {
		for n := 0; n < boardSize; n++ {
			if !g.tiles[i][n].hasValue() {
				return false
			}
		}
	}
	return true
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenHeight, ScreenWidth
}

func main() {
	ebiten.SetWindowSize(ScreenHeight, ScreenWidth)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(newGame()); err != nil {
		log.Fatal(err)
	}
}
