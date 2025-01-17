package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	tileSize    = 150
	tileMargin  = 160
	tileSpacing = 50
)

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(BACKGROUND_COLOR)

	if g.running {
		drawBoardLines(screen)
	}
}

func drawBoardLines(screen *ebiten.Image) {
	for i := 0; i <= boardWidth; i++ {
		for n := 0; n < boardHeight; n++ {
			line := ebiten.NewImage(3, 153)
			line.Fill(COLOR_LIGHT_GREY)
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(i*tileSize+tileSpacing), float64(n*tileSize+tileSpacing))
			screen.DrawImage(line, op)
		}
	}

	for i := 0; i < boardWidth; i++ {
		for n := 0; n <= boardHeight; n++ {
			line := ebiten.NewImage(150, 3)
			line.Fill(COLOR_LIGHT_GREY)
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(i*tileSize+tileSpacing), float64(n*tileSize+tileSpacing))
			screen.DrawImage(line, op)
		}
	}
}
