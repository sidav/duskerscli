package main

type room struct {
	name string
	// actorsHere
	// featuresHere
	isExplored bool
	isSeen     bool
	conns      [2]*connection // either to down or to right
}

type connection struct {
	rcx, rcy                     int // RelativeCoordX, RelativeCoordY. Can be either 0 or 1, NOT -1.
	isDoor                       bool
	lockLevel                    int
	isOpened, isBroken, isLocked bool
}
