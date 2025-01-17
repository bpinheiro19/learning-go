package main

import "image/color"

type Tile struct {
	letter Letter
}

type Letter struct {
	value byte
	color color.RGBA
}

func (t *Tile) isValid() bool {
	return t.letter.value >= 65 && t.letter.value <= 90
}

func (t *Tile) setGreenColor() {
	t.letter.color = COLOR_GREEN
}

func (t *Tile) setYellowColor() {
	t.letter.color = COLOR_YELLOW
}

func (t *Tile) setGrayColor() {
	t.letter.color = COLOR_GREY
}
