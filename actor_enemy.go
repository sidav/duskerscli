package main

func (a *actor) enemyAct() {
	// temp
	// try to do damage or use ability
	switch a.staticId {
	case ACTOR_MUTANT:
		actsHere := CURR_LEVEL.getAllActorsAtCoords(a.x, a.y)
		for _, actor := range actsHere {
			if actor.isPlayerControlled && actor.hp > 0 {
				actor.hp -= 1
				CURR_LEVEL.appendToLogMessage("%s is attacked by %s!", actor.getName(), a.getName())
				a.spendTimeForAction(10)
				return
			}
		}
	}
	// move by random vector
	for try := 0; try < 10; try++ {
		vx, vy := auxrnd.RandomUnitVectorInt()
		if CURR_LEVEL.canActorMoveByVector(a, vx, vy) {
			CURR_LEVEL.moveActorByVector(a, vx, vy)
			a.spendTimeForAction(20)
			return
		}
	}
}
