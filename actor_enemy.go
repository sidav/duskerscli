package main

func (a *actor) enemy_act() {
	// temp
	// move by random vector
	for try := 0; try < 10; try++ {
		vx, vy := auxrnd.RandomUnitVectorInt()
		if CURR_LEVEL.canActorMoveByVector(a, vx, vy) {
			CURR_LEVEL.moveActorByVector(a, vx, vy)
			return
		}
	}
}
