package main

func initLevel() *level {
	lvl := &level{}
	const LEVELSIZE = 4
	lvl.rooms = make([][]*room, LEVELSIZE)
	for i := range(lvl.rooms) {
		lvl.rooms[i] = make([]*room, LEVELSIZE)
		for j := range lvl.rooms[i] {
			if !rnd.OneChanceFrom(0) {
				lvl.rooms[i][j] = &room{
					name:       "",
					isExplored: false,
					isSeen:     false,
					conns:      [2]*connection{},
				}
				if i < LEVELSIZE-1 && !rnd.OneChanceFrom(6) {
					lvl.rooms[i][j].conns[0] = &connection{
						rcx:       1,
						rcy:       0,
						isDoor:    true,
						lockLevel: 0,
						isOpened:  rnd.OneChanceFrom(4),
						isBroken:  false,
						isLocked:  false,
					}
				}
				if j < LEVELSIZE-1 && !rnd.OneChanceFrom(6) {
					lvl.rooms[i][j].conns[1] = &connection{
						rcx:       0,
						rcy:       1,
						isDoor:    true,
						lockLevel: 0,
						isOpened:  rnd.OneChanceFrom(4),
						isBroken:  false,
						isLocked:  false,
					}
				}
			}
		}
	}

	lvl.actors = append(lvl.actors, &actor{
		name: "alpha",
		staticId: ACTOR_DRONE,
		hp: 5,
		x:  0,
		y:  0,
		isPlayerControlled: true,
	})
	lvl.actors = append(lvl.actors, &actor{
		name: "bravo",
		staticId: ACTOR_DRONE,
		hp: 5,
		x:  0,
		y:  0,
		isPlayerControlled: true,
	})
	lvl.actors = append(lvl.actors, &actor{
		staticId: ACTOR_MUTANT,
		hp: 5,
		x:  3,
		y:  3,
		isPlayerControlled: false,
	})

	return lvl
}
