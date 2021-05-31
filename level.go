package main

import "fmt"

type level struct {
	rooms             [][]*room
	actors            []*actor
	currLog           string
	currentTurnNumber int
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

func (l *level) getActorByName(name string) *actor {
	for _, a := range l.actors {
		if a.name == name {
			return a
		}
	}
	return nil
}

func (l *level) setLogMessage(msg string, args ...interface{}) {
	l.currLog = fmt.Sprintf(msg, args...)
}

func (l *level) appendToLogMessage(msg string, args ...interface{}) {
	l.currLog += fmt.Sprintf(" "+msg, args...)
}

func (l *level) getConnBetweenRoomsAtCoords(x1, y1, x2, y2 int) *connection {
	diffX := x1 - x2
	diffY := y1 - y2
	dist := euclideanDistance(x1, y1, x2, y2)
	if dist == 0 || dist > 1 {
		return nil
	}
	neededRoom := l.rooms[x1][y1]
	if x1 > x2 || y1 > y2 {
		neededRoom = l.rooms[x2][y2]
	}
	for _, c := range neededRoom.conns {
		if c == nil {
			continue
		}
		if c.rcx == abs(diffX) && c.rcy == abs(diffY) {
			return c
		}
	}
	return nil
}
