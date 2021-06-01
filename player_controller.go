package main

import "strings"

type playerController struct {
	l *level
}

func (p *playerController) playerTurn() {
	input := rend.readPlayerInput()
	abortGame = input == "exit"
	p.parsePlayerInput(input)
}

func (p *playerController) parsePlayerInput(inputString string) {
	splitted := strings.Split(inputString, " ")
	// spLen := len(splitted)
	switch splitted[0] {
	case "move":
		actor := CURR_LEVEL.getActorByName(splitted[1])
		if actor == nil {
			return
		}
		trx, try := p.strToRoomCoords(splitted[2])
		if trx < 0 {
			return
		}
		actor.currOrder = &order{
			x:           trx,
			y:           try,
			tx:          trx,
			ty:          try,
			orderTypeId: ORDER_MOVE,
		}
		CURR_LEVEL.setLogMessage("Order set!")
	case "door":
		x1, y1, x2, y2 := p.strToTwoRoomsCoords(splitted[1])
		conn := CURR_LEVEL.getConnBetweenRoomsAtCoords(x1, y1, x2, y2)
		if conn != nil {
			conn.isOpened = !conn.isOpened
			CURR_LEVEL.appendToLogMessage("Opened.")
		} else {
			CURR_LEVEL.appendToLogMessage("Can't find conn at %d,%d - %d,%d.", x1, y1, x2, y2)
		}
	}
}

func (p *playerController) strToRoomCoords(s string) (int, int) {
	if len(s) != 2 {
		return -1, -1
	}
	return int(s[0] - "a"[0]), int(s[1] - "0"[0] - 1)
}

func (p *playerController) strToTwoRoomsCoords(s string) (int, int, int, int) {
	if len(s) != 4 {
		return -1, -1, -1, -1
	}
	x1, y1 := p.strToRoomCoords(s[:2])
	x2, y2 := p.strToRoomCoords(s[2:])
	return x1, y1, x2, y2
}
