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
