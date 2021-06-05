package main

type actor struct {
	name               string
	nextTurnToAct      int
	staticId           uint8
	hp                 int
	x, y               int // room-wise
	modules            []*module
	isPlayerControlled bool
	currOrder          *order
}

func (a *actor) getName() string {
	if a.name != "" {
		return a.name
	}
	return a.getStaticData().defaultName
}

func (a *actor) getStaticData() *actorStaticData {
	return staticDataTableActors[a.staticId]
}

func (a *actor) spendTimeForAction(t int) {
	a.nextTurnToAct = CURR_LEVEL.currentTurnNumber + t
}

func (a *actor) isTimeToAct() bool {
	return a.nextTurnToAct <= CURR_LEVEL.currentTurnNumber
}

func (a *actor) act() {
	if a.isPlayerControlled {
		a.executeOrder()
	} else {
		a.enemyAct()
	}
}

func (a *actor) executeOrder() {
	if a.currOrder == nil {
		return
	}
	switch a.currOrder.orderTypeId {
	case ORDER_MOVE:
		vx := a.currOrder.x - a.x
		vy := a.currOrder.y - a.y
		vx, vy = toUnitVector(vx, vy)
		if !CURR_LEVEL.canActorMoveByVector(a, vx, vy) {
			CURR_LEVEL.appendToLogMessage("Can't move from %d,%d to %d,%d", a.x, a.y, a.currOrder.x, a.currOrder.y)
			a.currOrder = nil
			return
		}
		CURR_LEVEL.moveActorByVector(a, vx, vy)
		if a.x == a.currOrder.x && a.y == a.currOrder.y {
			CURR_LEVEL.appendToLogMessage("%s: arrived.", a.name)
			a.currOrder = nil
			return
		}
	}
}

func (a *actor) acquireEnergy(amount int) {
	for _, m := range a.modules {
		if m.currentEnergyCharge + amount <= m.staticData.addsEnergyStorage {
			m.currentEnergyCharge += amount
			amount = 0
			break
		} else {
			amount -= m.staticData.addsEnergyStorage - m.currentEnergyCharge
			m.currentEnergyCharge = m.staticData.addsEnergyStorage
		}
	}
}

func (a *actor) spendEnergy(amount int) {
	for _, m := range a.modules {
		if m.currentEnergyCharge >= amount {
			m.currentEnergyCharge -= amount
			amount = 0
			break
		} else {
			amount -= m.currentEnergyCharge
			m.currentEnergyCharge = 0
		}
	}
	if amount > 0 {
		panic("ENERGY SPENDING ERROR")
	}
}

func (a *actor) getEnergyCurrAndMax() (int, int) {
	max := 0
	curr := 0
	for _, m := range a.modules {
		curr += m.currentEnergyCharge
		max += m.staticData.addsEnergyStorage
	}
	if curr > max {
		curr = max
		panic("ENERGY > MAX")
	}
	return curr, max
}

func (a *actor) applyAllModules() {
	for _, mod := range a.modules {
		ce, _ := a.getEnergyCurrAndMax()
		if !mod.staticData.activatable || mod.isEnabled {
			if ce >= mod.staticData.drainsEnergy {
				a.spendEnergy(mod.staticData.drainsEnergy)
				for _, eff := range mod.staticData.effects {
					mod.applyModuleEffect(a, eff)
				}
			} else {
				mod.isEnabled = false
				CURR_LEVEL.appendToLogMessage("%s: %s forcefully disabled!", a.name, mod.getName())
			}
		}
	}
}
