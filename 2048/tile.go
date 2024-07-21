package main

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
