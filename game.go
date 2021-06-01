package main

type game struct {}

func (g *game) gameLoop() {
	lvl := initLevel()
	CURR_LEVEL = lvl
	rend = initRenderer()
	p := playerController{}

	for !abortGame {
		lvl.resetRoomsVisibility()

		for _, a := range CURR_LEVEL.actors {
			for _, mod := range a.modules {
				if mod.isEnabled {
					for _, eff := range mod.staticData.effects {
						eff.applyModuleEffect(a)
					}
				}
			}
		}

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
