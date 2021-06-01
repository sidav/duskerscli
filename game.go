package main

type game struct {}

func (g *game) gameLoop() {
	lvl := initLevel()
	CURR_LEVEL = lvl
	rend = initRenderer()
	p := playerController{}

	for !abortGame {
		// lvl.setLogMessage("")

		rend.render(lvl)
		if lvl.currentTurnNumber % 10 == 0 {
			p.playerTurn()
		}

		for _, a := range CURR_LEVEL.actors {
			if !a.isTimeToAct() {
				continue
			}
			if a.isPlayerControlled {
				a.executeOrder()
			} else {
				a.enemy_act()
			}
		}

		lvl.currentTurnNumber++
	}
}
