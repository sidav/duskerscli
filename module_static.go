package main

const (
	MODULE_BATTERY = iota
	MODULE_SCANNER = iota
)

type moduleStaticData struct {
	activatable       bool
	defaultName       string
	addsEnergyStorage int
	drainsEnergy      int
	effects           []*moduleEffect
}

var staticModuleDataTable = map[int]*moduleStaticData{
	MODULE_BATTERY: {
		defaultName:       "Battery",
		addsEnergyStorage: 100,
		drainsEnergy:      0,
	},
	MODULE_SCANNER: {
		activatable:       true,
		defaultName:       "Scanner",
		addsEnergyStorage: 0,
		drainsEnergy:      5,
		effects: []*moduleEffect{
			{code: EFFECT_SURVEY},
		},
	},
}
