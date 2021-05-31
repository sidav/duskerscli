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
		if trx*try < 0 {
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
	return int("a"[0] - s[0]), int("0"[0] - s[1])
}
