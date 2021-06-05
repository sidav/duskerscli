package main

const (
	FACILITY_INTERFACE uint8 = iota
	FACILITY_GENERATOR
	FACILITY_TURRET
)

type roomFacility struct {
	code        uint8
	number      int
	hp          int
	destroyable bool
	isActive    bool
}

func (rf *roomFacility) getAppearanceChar() rune {
	switch rf.code {
	case FACILITY_INTERFACE:
		return '#'
	case FACILITY_TURRET:
		return 'T'
	case FACILITY_GENERATOR:
		return '%'
	}
	panic("No rune for facility.")
}
