package main

const (
	MODULE_BATTERY = iota
	MODULE_EMERGENCY_GENERATOR
	MODULE_SCANNER
)

type moduleStaticData struct {
	activatable                 bool
	immobilizesActorWhileActive bool

	defaultName string

	addsEnergyStorage int
	drainsEnergy      int

	effects []*moduleEffect
}

var staticModuleDataTable = map[int]*moduleStaticData{
	MODULE_BATTERY: {
		defaultName:       "Battery",
		addsEnergyStorage: 100,
		drainsEnergy:      0,
	},
	MODULE_EMERGENCY_GENERATOR: {
		defaultName:                 "Emerg.generator",
		activatable:                 true,
		immobilizesActorWhileActive: true,
		addsEnergyStorage:           0,
		drainsEnergy:                0,
		effects: []*moduleEffect{
			{EFFECT_GENERATE_ENERGY},
		},
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
