package main

const (
	MODULE_BATTERY = iota
	MODULE_SCANNER = iota
)

type moduleStaticData struct {
	defaultName       string
	addsEnergyStorage int
	drainsEnergy      int
}

var staticModuleDataTable = map[int]*moduleStaticData {
	MODULE_BATTERY: {
		defaultName:       "Battery",
		addsEnergyStorage: 100,
		drainsEnergy:      0,
	},
	MODULE_SCANNER: {
		defaultName:       "Scanner",
		addsEnergyStorage: 0,
		drainsEnergy:      5,
	},
}
