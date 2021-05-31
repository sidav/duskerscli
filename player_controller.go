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
	}
}

func (p *playerController) strToRoomCoords(s string) (int, int) {
	if len(s) != 2 {
		return -1, -1
	}
	return int(s[0] - "a"[0]), int(s[1] - "0"[0])
}
