package main

import "duskerscli/console_wrapper"

func main() {
	lvl := initLevel()
	rend := initRenderer()
	rend.render(lvl)
	console_wrapper.ReadKey()
	console_wrapper.Close_console()
}
