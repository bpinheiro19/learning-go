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
	tiles    [][]Tile
	running  bool
	gameover bool
	win      bool
}

const (
	ScreenWidth  = 700
	ScreenHeight = 700
	boardSize    = 4

	tileSize    = 150
	tileMargin  = 160
	tileSpacing = 50
)

var (
	mplusFaceSource *text.GoTextFaceSource
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

	return !g.CanMoveUp() && !g.CanMoveDown() && !g.CanMoveRight() && !g.CanMoveLeft()
}

func (g *Game) CanMoveUp() bool {
	for n := 0; n < boardSize; n++ {
		i := 0
		for i < 4 {
			val := g.tiles[i][n].value
			if val != 0 && i-1 >= 0 {
				if g.tiles[i][n].value == g.tiles[i-1][n].value {
					return true
				}
			}
			i++
		}
	}
	return false
}

func (g *Game) CanMoveDown() bool {
	for n := 0; n < boardSize; n++ {
		i := 3
		for i >= 0 {
			val := g.tiles[i][n].value
			if val != 0 && i+1 < 4 {
				if g.tiles[i][n].value == g.tiles[i+1][n].value {
					return true
				}
			}
			i--
		}
	}
	return false
}

func (g *Game) CanMoveRight() bool {
	for i := 0; i < boardSize; i++ {
		n := 3
		for n >= 0 {
			val := g.tiles[i][n].value
			if val != 0 && n+1 < 4 {
				if g.tiles[i][n].value == g.tiles[i][n+1].value {
					return true
				}
			}
			n--
		}
	}
	return false
}

func (g *Game) CanMoveLeft() bool {
	for i := 0; i < boardSize; i++ {
		n := 0
		for n < 4 {
			val := g.tiles[i][n].value
			if val != 0 && n-1 >= 0 {
				if g.tiles[i][n].value == g.tiles[i][n-1].value {
					return true
				}
			}
			n++
		}
	}
	return false
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
						g.tiles[i][n-1].value = 2 * g.tiles[i][n].value
						g.tiles[i][n].value = 0
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

	if g.running {

		if g.gameover || g.win {
			g.running = false
		}

		if repeatingKeyPressed(ebiten.KeyUp) {
			g.moveTilesUp()
			g.spawnTile()
			g.printBoard()
			g.gameover = g.checkGameOver()
			g.win = g.checkWin()
		}

		if repeatingKeyPressed(ebiten.KeyDown) {
			g.moveTilesDown()
			g.spawnTile()
			g.printBoard()
			g.gameover = g.checkGameOver()
			g.win = g.checkWin()
		}

		if repeatingKeyPressed(ebiten.KeyLeft) {
			g.moveTilesLeft()
			g.spawnTile()
			g.printBoard()
			g.gameover = g.checkGameOver()
			g.win = g.checkWin()
		}

		if repeatingKeyPressed(ebiten.KeyRight) {
			g.moveTilesRight()
			g.spawnTile()
			g.printBoard()
			g.gameover = g.checkGameOver()
			g.win = g.checkWin()
		}

	} else {
		//Restart Game
		if repeatingKeyPressed(ebiten.KeyEnter) {
			g.initBoard()
			g.running = true
			g.gameover = false
			g.win = false
		}
	}

	return nil
}

func DrawBoardLines(screen *ebiten.Image) {
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
}

func (g *Game) DrawTiles(screen *ebiten.Image) {
	for i := 0; i < boardSize; i++ {
		for n := 0; n < boardSize; n++ {
			tile := g.tiles[i][n]
			if tile.value != 0 {

				col, fontsize, x, y := tile.GetTileConfig()

				tileImage := ebiten.NewImage(tileSize-4, tileSize-4)
				tileImage.Fill(col)
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(n*tileSize+tileSpacing+3), float64(i*tileSize+tileSpacing+3))
				screen.DrawImage(tileImage, op)

				optext := &text.DrawOptions{}

				optext.GeoM.Translate(float64(n*tileSize+tileSpacing+x), float64(i*tileSize+tileSpacing+y))
				optext.ColorScale.ScaleWithColor(color.Black)
				text.Draw(screen, strconv.Itoa(tile.value), &text.GoTextFace{
					Source: mplusFaceSource,
					Size:   float64(fontsize),
				}, optext)
			}
		}
	}
}

func DrawGameOverScreen(screen *ebiten.Image) {
	optext := &text.DrawOptions{}

	optext.GeoM.Translate(float64(120), float64(200))
	optext.ColorScale.ScaleWithColor(color.Black)
	text.Draw(screen, "Game Over", &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   float64(90),
	}, optext)

	optext.GeoM.Translate(float64(35), float64(200))
	text.Draw(screen, "Press Enter to restart the game", &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   float64(25),
	}, optext)
}

func DrawWinScreen(screen *ebiten.Image) {
	optext := &text.DrawOptions{}

	optext.GeoM.Translate(float64(160), float64(200))
	optext.ColorScale.ScaleWithColor(color.Black)
	text.Draw(screen, "You Win!", &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   float64(90),
	}, optext)

	optext.GeoM.Translate(float64(0), float64(200))
	text.Draw(screen, "Press Enter to restart the game", &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   float64(25),
	}, optext)
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(BACKGROUND_COLOR)

	if g.running {
		DrawBoardLines(screen)
		g.DrawTiles(screen)

	} else if g.win {
		DrawWinScreen(screen)

	} else if g.gameover {
		DrawGameOverScreen(screen)
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
