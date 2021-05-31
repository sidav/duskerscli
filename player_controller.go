package main

import cw "duskerscli/console_wrapper"

type playerController struct {}

func (p *playerController) readPlayerInput() {
	currLine := ""
	key := ""
	for key != "ESCAPE" {
		currLine, key = cw.ReadTextInputAndKeyPress("> ", currLine, 0, 17)
		if key == "ENTER" {
			currLine = ""
		}
	}
}
