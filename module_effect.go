package main

const (
	EFFECT_SURVEY uint8 = iota
	EFFECT_GENERATE_ENERGY

	EFFECT_SIMPLE_ATTACK
)

type moduleEffect struct {
	code   uint8
	damage int
}

func (me *moduleEffect) applyModuleEffect(user *actor) {
	switch me.code {
	case EFFECT_SURVEY:
		ux, uy := user.x, user.y
		for x := -1; x <= 1; x++ {
			for y := -1; y <= 1; y++ {
				room := CURR_LEVEL.getRoomAtCoords(ux+x, uy+y)
				if room == nil {
					continue
				}
				room.isSeen = true
				room.isExplored = true
			}
		}
	case EFFECT_GENERATE_ENERGY:
		user.acquireEnergy(1)
	case EFFECT_SIMPLE_ATTACK:
		actrs := CURR_LEVEL.getAllActorsAtCoords(user.x, user.y)
		for _, a := range actrs {
			if !a.isPlayerControlled {
				a.hp -= me.damage
				CURR_LEVEL.appendToLogMessage("%s: opening fire at %s. ", user.name, a.name)
				break
			}
		}
	}
}
