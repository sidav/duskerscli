package main

import (
	"duskerscli/console_wrapper"
	"duskerscli/fibrandom"
)

var (
	rend        *renderer
	rnd, auxrnd fibrandom.FibRandom
	abortGame   bool
	CURR_LEVEL  *level
)

func main() {
	rnd.InitDefault()
	auxrnd.InitDefault()

	g := game{}
	g.gameLoop()

	defer console_wrapper.Close_console()
}
