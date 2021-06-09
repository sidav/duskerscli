package main

const (
	EFFECT_SURVEY uint8 = iota
	EFFECT_MOTION_SCANNER
	EFFECT_GENERATE_ENERGY
	EFFECT_NETWORK_CONNECTION

	EFFECT_SIMPLE_ATTACK
)

type moduleEffect struct {
	code   uint8
	damage int
}

func (m *module) applyModuleEffect(user *actor, me *moduleEffect) {
	switch me.code {
	case EFFECT_SURVEY:
		ux, uy := user.x, user.y
		for x := -1; x <= 1; x++ {
			for y := -1; y <= 1; y++ {
				room := CURR_LEVEL.getRoomAtCoords(ux+x, uy+y)
				if room == nil {
					continue
				}
				room.isExplored = true
			}
		}
	case EFFECT_MOTION_SCANNER:
		ux, uy := user.x, user.y
		for x := -1; x <= 1; x++ {
			for y := -1; y <= 1; y++ {
				room := CURR_LEVEL.getRoomAtCoords(ux+x, uy+y)
				if room == nil {
					continue
				}
				room.isUnderMotionScanner = true
			}
		}
	case EFFECT_GENERATE_ENERGY:
		user.acquireEnergy(1)
	case EFFECT_NETWORK_CONNECTION:
		if CURR_LEVEL.hasFacilityAtCoords(FACILITY_INTERFACE, user.x, user.y) {
			CURR_LEVEL.showNetworkTerminal(user.x, user.y)
		} else {
			CURR_LEVEL.appendToLogMessage("But this room has no network interface!")
		}
		m.isEnabled = false
	case EFFECT_SIMPLE_ATTACK:
		actrs := CURR_LEVEL.getAllActorsAtCoords(user.x, user.y)
		for _, target := range actrs {
			if !target.isPlayerControlled && target.asFacility == nil {
				target.hp -= me.damage
				CURR_LEVEL.appendToLogMessage("%s: opening fire at %s. ", user.getName(), target.getName())
				break
			}
		}
	}
}
