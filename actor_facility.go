package main

const (
	FACILITY_INTERFACE uint8 = iota
	FACILITY_GENERATOR
	FACILITY_TURRET
)

type facility struct {
	code             uint8
	associatedNumber int
	destroyable      bool
	isActive         bool
}

func (rf *facility) getAppearanceChar() rune {
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
