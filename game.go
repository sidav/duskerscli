package main

type game struct {}

func (g *game) gameLoop() {
	lvl := initLevel()
	CURR_LEVEL = lvl
	rend = initRenderer()
	p := playerController{}

	for !abortGame {
		lvl.resetRoomsVisibility()

		if lvl.currentTurnNumber % 10 == 0 { // TODO: think of a better solution
			for _, a := range CURR_LEVEL.actors {
				a.applyAllModules()
			}
		}
		lvl.clearDestroyedActors()
		if lvl.currentTurnNumber % 10 == 0 {
			rend.render(lvl)
			lvl.setLogMessage("")
			p.playerTurn()
		}

		for _, a := range CURR_LEVEL.actors {
			if !a.isTimeToAct() {
				continue
			}
			if a.isPlayerControlled {
				a.executeOrder()
			} else {
				a.enemyAct()
			}
		}

		lvl.currentTurnNumber++
	}
}
