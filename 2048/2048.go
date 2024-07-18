package main

import (
	"bytes"
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	tiles []Tile
}

type Tile struct {
	value int
	x     int
	y     int
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
	BACKGROUND_COLOR = color.RGBA{207, 195, 176, 0xff}
	COLOR2           = color.RGBA{245, 230, 200, 0xff}
	COLOR4           = color.RGBA{245, 220, 175, 0xff}
	COLOR8           = color.RGBA{245, 210, 160, 0xff}
	COLOR16          = color.RGBA{245, 200, 145, 0xff}
	COLOR32          = color.RGBA{245, 190, 130, 0xff}
	COLOR64          = color.RGBA{245, 180, 115, 0xff}
	COLOR128         = color.RGBA{245, 140, 100, 0xff}
	COLOR256         = color.RGBA{245, 130, 85, 0xff}
	COLOR512         = color.RGBA{245, 110, 70, 0xff}
	COLOR1024        = color.RGBA{245, 80, 55, 0xff}
	COLOR2048        = color.RGBA{245, 60, 40, 0xff}
)

func (g *Game) Update() error {
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
			if g.tiles[i+n].value != 0 {
				tile := ebiten.NewImage(tileSize-4, tileSize-4)
				tile.Fill(COLOR2)
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(i*tileSize+tileSpacing+3), float64(n*tileSize+tileSpacing+3))
				screen.DrawImage(tile, op)

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
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
