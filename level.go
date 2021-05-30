package main

type level struct {
	rooms [][]*room
	actors []*actor
}

func (l *level) getAllActorsAtCoords(x, y int) []*actor {
	var actsHere []*actor
	for _, a := range l.actors {
		if a.x == x && a.y == y {
			actsHere = append(actsHere, a)
		}
	}
	return actsHere
}
