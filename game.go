package main

type game struct {}

func (g *game) gameLoop() {
	lvl := initLevel()
	CURR_LEVEL = lvl
	rend = initRenderer()
	for !abortGame {
		for _, a := range CURR_LEVEL.actors {
			a.executeOrder()
		}
		rend.render(lvl)
		p := playerController{}
		p.playerTurn()
		lvl.currentTurnNumber++
	}
}
