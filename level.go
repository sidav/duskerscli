package main

import (
	"duskerscli/pathfinder"
	"fmt"
)

type level struct {
	rooms             [][]*room
	actors            []*actor
	currLog           []string
	currentTurnNumber int

	pathfinder *pathfinder.AStarPathfinder
}

func (l *level) getRoomAtCoords(x, y int) *room {
	if x >= 0 && x < len(l.rooms) && y >= 0 && y < len(l.rooms[x]) {
		return l.rooms[x][y]
	}
	return nil
}

func (l *level) resetRoomsVisibility() {
	for x := range l.rooms {
		for y := range l.rooms[x] {
			if l.rooms[x][y] != nil {
				l.rooms[x][y].isSeenRightNow = false
				l.rooms[x][y].isUnderMotionScanner = false
			}
		}
	}
	for _, a := range l.actors {
		if a.isPlayerControlled {
			l.rooms[a.x][a.y].isExplored = true
			l.rooms[a.x][a.y].isSeenRightNow = true
		}
	}
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
		if stringBeginsWith(a.name, name) {
			if foundActor != nil {
				return nil // partial name belongs to more that one actor
			}
			foundActor = a
		}
	}
	return foundActor
}

func (l *level) getFacilitiesAt(x, y int) []*actor {
	var facsHere []*actor
	for _, a := range l.actors {
		if a.x == x && a.y == y && a.asFacility != nil {
			facsHere = append(facsHere, a)
		}
	}
	return facsHere
}

func (l *level) hasFacilityAtCoords(facCode uint8, x, y int) bool {
	for _, f := range l.actors {
		if f.asFacility != nil && f.asFacility.code == facCode && f.x == x && f.y == y {
			return true
		}
	}
	return false
}

func (l *level) setLogMessage(msg string, args ...interface{}) {
	l.currLog = append(l.currLog[1:], fmt.Sprintf(msg, args...))
}

func (l *level) appendToLogMessage(msg string, args ...interface{}) {
	l.currLog = append(l.currLog[1:], fmt.Sprintf(msg, args...))
	// l.currLog += fmt.Sprintf(" "+msg, args...)
}

func (l *level) canActorMoveByVector(a *actor, vx, vy int) bool {
	conn := l.getConnFromRoomByVector(a.x, a.y, vx, vy)
	return conn != nil && conn.isOpened
}

func (l *level) moveActorByVector(a *actor, vx, vy int) {
	if l.canActorMoveByVector(a, vx, vy) {
		a.x += vx
		a.y += vy
	}
	a.spendTimeForAction(10)
}

func (l *level) clearDestroyedActors() {
	for i := len(l.actors) - 1; i >= 0; i-- {
		if l.actors[i].hp <= 0 {
			l.appendToLogMessage("%s destroyed.", l.actors[i].getName())
			l.actors = append(l.actors[:i], l.actors[i+1:]...)
		}
	}
}

func (l *level) getAllConnectionsOfRoom(x, y int) []*connection {
	var conns []*connection
	conns = append(conns, l.getConnFromRoomByVector(x, y, 1, 0))
	conns = append(conns, l.getConnFromRoomByVector(x, y, -1, 0))
	conns = append(conns, l.getConnFromRoomByVector(x, y, 0, 1))
	conns = append(conns, l.getConnFromRoomByVector(x, y, 0, -1))
	return conns
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
