package main

import "fmt"

type level struct {
	rooms  [][]*room
	actors []*actor
	currLog    string
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
