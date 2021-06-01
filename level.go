package main

import (
	"fmt"
	"strings"
)

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
	var foundActor *actor
	for _, a := range l.actors {
		if strings.Index(a.name, name) == 0 {
			if foundActor != nil {
				return nil // partial name belongs to more that one actor
			}
			foundActor = a
		}
	}
	return foundActor
}

func (l *level) setLogMessage(msg string, args ...interface{}) {
	l.currLog = fmt.Sprintf(msg, args...)
}

func (l *level) appendToLogMessage(msg string, args ...interface{}) {
	l.currLog += fmt.Sprintf(" "+msg, args...)
}

func (l *level) canActorMoveByVector(a *actor, vx, vy int) bool {
	return l.getConnFromRoomByVector(a.x, a.y, vx, vy) != nil
}

func (l *level) moveActorByVector(a *actor, vx, vy int) {
	if l.canActorMoveByVector(a, vx, vy) {
		a.x += vx
		a.y += vy
	}
}

func (l *level) getConnFromRoomByVector(x, y, vx, vy int) *connection {
	if vx < 0 {
		vx = -vx
		x--
	}
	if vy < 0 {
		vy = -vy
		y--
	}
	if x < 0 || x >= len(l.rooms) || y < 0 || y >= len(l.rooms[0]) {
		return nil
	}
	room := l.rooms[x][y]
	if room == nil {
		return nil
	}
	for _, c := range room.conns {
		if c == nil {
			continue
		}
		if c.rcx == vx && c.rcy == abs(vy) {
			return c
		}
	}
	return nil
}

func (l *level) getConnBetweenRoomsAtCoords(x1, y1, x2, y2 int) *connection {
	return l.getConnFromRoomByVector(x1, y1, x2-x1, y2-y1)
}
