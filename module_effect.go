package main

const (
	EFFECT_SURVEY uint8 = iota
)

type moduleEffect struct {
	code uint8
}

func (me *moduleEffect) applyModuleEffect(user *actor) {
	switch me.code {
	case EFFECT_SURVEY:
		ux, uy := user.x, user.y
		for x := -1; x <= 1; x++ {
			for y := -1; y <= 1; y++ {
				CURR_LEVEL.rooms[ux+x][uy+y].isSeen = true
				CURR_LEVEL.rooms[ux+x][uy+y].isExplored = true
			}
		}
	}
}
