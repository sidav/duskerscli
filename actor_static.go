package main

const (
	ACTOR_DRONE uint8 = iota
	ACTOR_MUTANT
)

type actorStaticData struct {
	char        rune
	defaultName string
	maxHp       int
}

var staticDataTableActors = map[uint8]*actorStaticData{
	ACTOR_DRONE: {
		char:        '@',
		defaultName: "Drone",
		maxHp:       10,
	},
	ACTOR_MUTANT: {
		char:        'm',
		defaultName: "Mutant",
		maxHp:       10,
	},
}
