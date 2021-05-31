package main

import cw "duskerscli/console_wrapper"

type playerController struct {}

func (p *playerController) playerTurn() {
	input := p.readPlayerInput()
	abortGame = input == "exit"
}

func (p *playerController) readPlayerInput() string {
	currLine := ""
	key := ""
	for key != "ENTER" {
		currLine, key = cw.ReadTextInputAndKeyPress("> ", currLine, 0, 17)
	}
	return currLine
}
