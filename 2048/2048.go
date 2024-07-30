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
			fmt.Print(g.tiles[n][i].value)

		}
		fmt.Println(" | ")
	}
	fmt.Println(" ----------------- ")
}

func (g *Game) moveTile(x, y, w, z int) {
	val := g.tiles[x][y].value
	g.tiles[x][y].value = 0
	g.tiles[w][z].value = val
}

func (g *Game) spawnTile() {
	availableZero := false

	for i := 0; i < boardSize; i++ {
		for n := 0; n < boardSize; n++ {
			if g.tiles[i][n].value == 0 {
				availableZero = true
			}
		}
	}

	for availableZero {

		x := rand.Intn(4)
		y := rand.Intn(4)

		val := func() int {
			if rand.Intn(10) < 9 {
				return 2
			}
			return 4
		}()

		if g.tiles[x][y].value == 0 {
			availableZero = false
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

func (g *Game) moveTilesUp() {
	for n := 0; n < boardSize; n++ {
		i := 0
		for i < 4 {
			val := g.tiles[i][n].value
			if val != 0 && i-1 >= 0 {

				if !(g.tiles[i-1][n].hasValue()) {
					g.moveTile(i, n, i-1, n)
					i -= 2
				} else {
					if g.tiles[i][n].value == g.tiles[i-1][n].value {
						g.tiles[i-1][n].value = 2 * g.tiles[i][n].value
						g.tiles[i][n].value = 0
					}
				}

			}
			i++
		}
	}
}

func (g *Game) moveTilesDown() {
	for n := 0; n < boardSize; n++ {
		i := 3
		for i >= 0 {
			val := g.tiles[i][n].value
			if val != 0 && i+1 < 4 {

				if !(g.tiles[i+1][n].hasValue()) {
					g.moveTile(i, n, i+1, n)
					i -= 2
				} else {
					if g.tiles[i][n].value == g.tiles[i+1][n].value {
						g.tiles[i+1][n].value = 2 * g.tiles[i][n].value
						g.tiles[i][n].value = 0
					}
				}

			}
			i--
		}
	}
}

func (g *Game) moveTilesRight() {
	for i := 0; i < boardSize; i++ {
		n := 3
		for n >= 0 {
			val := g.tiles[i][n].value
			if val != 0 && n+1 < 4 {
				if !(g.tiles[i][n+1].hasValue()) {
					g.moveTile(i, n, i, n+1)
					n += 2
				} else {
					if g.tiles[i][n].value == g.tiles[i][n+1].value {
						g.tiles[i][n+1].value = 2 * g.tiles[i][n].value
						g.tiles[i][n].value = 0
					}
				}

			}
			n--
		}
	}
}

func (g *Game) moveTilesLeft() {
	for i := 0; i < boardSize; i++ {
		n := 0
		for n < 4 {
			val := g.tiles[i][n].value
			if val != 0 && n-1 >= 0 {

				if !(g.tiles[i][n-1].hasValue()) {
					fmt.Println(i, n, i, n-1)
					g.moveTile(i, n, i, n-1)
					n -= 2
				} else {
					if g.tiles[i][n].value == g.tiles[i][n-1].value {
						fmt.Println(g.tiles[i][n].value)
						g.tiles[i][n-1].value = 2 * g.tiles[i][n].value
						g.tiles[i][n].value = 0
						fmt.Println(g.tiles[i][n-1].value)
					}
				}

			}
			n++
		}
	}
}

func repeatingKeyPressed(key ebiten.Key) bool {
	d := inpututil.KeyPressDuration(key)
	return d >= 1 && d < 2
}

func (g *Game) Update() error {

	if repeatingKeyPressed(ebiten.KeyUp) {
		g.spawnTile()
		g.printBoard()
	}

	if repeatingKeyPressed(ebiten.KeyDown) {
		g.spawnTile()
		g.printBoard()
	}

	if repeatingKeyPressed(ebiten.KeyLeft) {
		g.spawnTile()
		g.printBoard()
	}

	if repeatingKeyPressed(ebiten.KeyRight) {
		g.spawnTile()
		g.printBoard()
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
