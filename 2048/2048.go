package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	board []Tile
}

type Tile struct {
	value int
	x     int
	y     int
}

const (
	ScreenWidth  = 600
	ScreenHeight = 600
	boardSize    = 4

	tileSize   = 55
	tileMargin = 4
)

var (
	backgroundColor = color.RGBA{0xbb, 0xad, 0xa0, 0xff}
)

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(backgroundColor)

	for n := 0; n < 9; n++ {
		for i := 0; i < 9; i++ {
			tile := ebiten.NewImage(tileSize, tileSize)
			tile.Fill(color.White)
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(i*60+30), float64(n*60+10))
			screen.DrawImage(tile, op)
		}
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenHeight, ScreenWidth
}

func main() {
	ebiten.SetWindowSize(ScreenHeight, ScreenWidth)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
