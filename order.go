package main

const (
	ORDER_MOVE uint8 = iota
)

type order struct {
	x, y   int
	tx, ty int
	orderTypeId uint8
}
