package main

import "fmt"

type module struct {
	isEnabled           bool
	currentEnergyCharge int // for batteries
	staticData          *moduleStaticData
}

func (m *module) getName() string {
	return m.staticData.defaultName
}

func (m *module) getNameAndEnabled() string {
	enStr := ": up"
	if m.staticData.addsEnergyStorage > 0 {
		enStr = fmt.Sprintf(": E %d/%d", m.currentEnergyCharge, m.staticData.addsEnergyStorage)
	}
	if m.staticData.activatable {
		enStr = ": off"
		if m.isEnabled {
			enStr = ": on"
		}
	}
	return m.staticData.defaultName + enStr
}

func createModuleByStaticCode(code int) *module {
	mod := &module{
		isEnabled:  false,
		staticData: staticModuleDataTable[code],
	}
	mod.currentEnergyCharge = mod.staticData.addsEnergyStorage
	return mod
}
