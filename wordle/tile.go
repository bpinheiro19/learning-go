package main

import "image/color"

type Tile struct {
	letter Letter
	color  color.RGBA
}

type Letter struct {
	value byte
}

func (t *Tile) isValid() bool {
	return t.letter.value >= 65 && t.letter.value <= 90
}

func (t *Tile) setGreenColor() {
	t.color = COLOR_GREEN
}

func (t *Tile) setYellowColor() {
	t.color = COLOR_YELLOW
}

func (t *Tile) setGrayColor() {
	t.color = COLOR_GREY
}
