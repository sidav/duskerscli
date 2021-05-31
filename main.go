package main

import (
	"duskerscli/console_wrapper"
	"duskerscli/fibrandom"
)

var (
	rnd, auxrnd fibrandom.FibRandom
	abortGame bool
)

func main() {
	rnd.InitDefault()
	auxrnd.InitDefault()

	g := game{}
	g.gameLoop()

	defer console_wrapper.Close_console()
}
