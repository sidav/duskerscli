package main

type actor struct {
	name               string
	staticId           uint8
	hp                 int
	x, y               int // room-wise
	isPlayerControlled bool
	currOrder          *order
}

func (a *actor) getStaticData() *actorStaticData {
	return staticDataTableActors[a.staticId]
}

func (a *actor) act() {
	if a.isPlayerControlled {
		a.executeOrder()
	} else {
		a.enemy_act()
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
		if a.x == a.currOrder.x && a.y == a.currOrder.y {
			CURR_LEVEL.appendToLogMessage("%s: arrived.", a.name)
			a.currOrder = nil
			return
		}
		if !CURR_LEVEL.canActorMoveByVector(a, vx, vy) {
			CURR_LEVEL.appendToLogMessage("Can't move from %d,%d to %d,%d", a.x, a.y, a.currOrder.x, a.currOrder.y)
			a.currOrder = nil
			return
		}
		CURR_LEVEL.moveActorByVector(a, vx, vy)
	}
}
