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

func (a *actor) executeOrder() {
	if a.currOrder == nil {
		return
	}
	switch a.currOrder.orderTypeId {
	case ORDER_MOVE:
		conn := CURR_LEVEL.getConnBetweenRoomsAtCoords(a.x, a.y, a.currOrder.x, a.currOrder.y)
		if conn == nil {
			CURR_LEVEL.setLogMessage("Can't move from %d,%d to %d,%d", a.x, a.y, a.currOrder.x, a.currOrder.y)
			a.currOrder = nil
			return
		}
		a.x = a.currOrder.x
		a.y = a.currOrder.y
		a.currOrder = nil
	}
}
