package main

type game struct {}

func (g *game) gameLoop() {
	lvl := initLevel()
	CURR_LEVEL = lvl
	rend = initRenderer()
	for !abortGame {
		rend.render(lvl)
		p := playerController{}
		p.playerTurn()
	}
}
