package main

type game struct {}

func (g *game) gameLoop() {
	lvl := initLevel()
	rend := initRenderer()
	for !abortGame {
		rend.render(lvl)
		p := playerController{}
		p.playerTurn()
	}
}
