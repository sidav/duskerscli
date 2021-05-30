package main

type actor struct {
	staticId           uint8
	hp                 int
	x, y               int // room-wise
	isPlayerControlled bool
}

func (a *actor) getStaticData() *actorStaticData {
	return staticDataTableActors[a.staticId]
}