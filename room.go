package main

type room struct {
	name string
	// actorsHere
	facilitiesHere       []*roomFacility
	isExplored           bool
	isSeenRightNow       bool
	isUnderMotionScanner bool
	conns                [2]*connection // either to down or to right
}

type connection struct {
	rcx, rcy                     int // RelativeCoordX, RelativeCoordY. Can be either 0 or 1, NOT -1.
	isDoor                       bool
	lockLevel                    int
	isOpened, isBroken, isLocked bool
}

func (r *room) hasFacility(facCode uint8) bool {
	for _, f := range r.facilitiesHere {
		if f.code == facCode {
			return true
		}
	}
	return false
}
