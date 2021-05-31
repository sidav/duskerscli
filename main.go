package main

import "duskerscli/console_wrapper"

func main() {
	lvl := initLevel()
	rend := initRenderer()
	rend.render(lvl)
	p := playerController{}
	p.readPlayerInput()

	defer console_wrapper.Close_console()
}
