package main

import "image/color"

type Tile struct {
	value int
	x     int
	y     int
}

func (t *Tile) getPos() (int, int) {
	return t.x, t.y
}

func (t *Tile) hasValue() bool {
	return t.value != 0
}

func (t *Tile) hasMaxValue() bool {
	return t.value == 2048
}

func (t *Tile) GetTileConfig() (color.RGBA, int, int, int) {
	switch t.value {
	case 4:
		return COLOR4, 90, 45, 15
	case 8:
		return COLOR8, 90, 45, 15
	case 16:
		return COLOR16, 80, 25, 20
	case 32:
		return COLOR32, 80, 25, 20
	case 64:
		return COLOR64, 80, 25, 20
	case 128:
		return COLOR128, 70, 10, 25
	case 256:
		return COLOR256, 70, 10, 25
	case 512:
		return COLOR512, 70, 10, 25
	case 1024:
		return COLOR1024, 55, 5, 35
	case 2048:
		return COLOR2048, 55, 5, 35
	}
	return COLOR2, 90, 45, 15
}
