package main

import "duskerscli/pathfinder"

func initLevel() *level {
	lvl := &level{}
	lvl.pathfinder = &pathfinder.AStarPathfinder{
		DiagonalMoveAllowed:       false,
		ForceGetPath:              true,
		ForceIncludeFinish:        false,
		AutoAdjustDefaultMaxSteps: false,
	}
	const LEVELSIZE = 4
	lvl.rooms = make([][]*room, LEVELSIZE)
	for i := range(lvl.rooms) {
		lvl.rooms[i] = make([]*room, LEVELSIZE)
		for j := range lvl.rooms[i] {
			if !rnd.OneChanceFrom(8) {
				lvl.rooms[i][j] = &room{
					name:           "",
					isExplored:     false,
					isSeenRightNow: false,
					conns:          [2]*connection{},
				}
			}
		}
	}

	for x := range lvl.rooms {
		for y := range lvl.rooms[x] {
			if lvl.rooms[x][y] != nil {
				if x < LEVELSIZE-1 && lvl.rooms[x+1][y] != nil && !rnd.OneChanceFrom(5) {
					lvl.rooms[x][y].conns[0] = &connection{
						rcx:       1,
						rcy:       0,
						isDoor:    true,
						lockLevel: 0,
						isOpened:  rnd.OneChanceFrom(2),
						isBroken:  false,
						isLocked:  false,
					}
				}
				if y < LEVELSIZE-1 && lvl.rooms[x][y+1] != nil && !rnd.OneChanceFrom(5) {
					lvl.rooms[x][y].conns[1] = &connection{
						rcx:       0,
						rcy:       1,
						isDoor:    true,
						lockLevel: 0,
						isOpened:  rnd.OneChanceFrom(2),
						isBroken:  false,
						isLocked:  false,
					}
				}
			}
		}
	}

	startRoom, srx, sry := lvl.getRandomRoomInRange(0, 0, 1, 1)

	lvl.actors = append(lvl.actors, &actor{
		name: "Alpha",
		staticId: ACTOR_DRONE,
		hp: 5,
		x:  srx,
		y:  sry,
		isPlayerControlled: true,
		modules: []*module{
			createModuleByStaticCode(MODULE_BATTERY),
			createModuleByStaticCode(MODULE_EMERGENCY_GENERATOR),
			createModuleByStaticCode(MODULE_SURVEYOR),
		},
	})
	lvl.actors = append(lvl.actors, &actor{
		name: "Bravo",
		staticId: ACTOR_DRONE,
		hp: 5,
		x:  srx,
		y:  sry,
		isPlayerControlled: true,
		modules: []*module{
			createModuleByStaticCode(MODULE_BATTERY),
			createModuleByStaticCode(MODULE_BATTERY),
			createModuleByStaticCode(MODULE_NETWORKING),
			createModuleByStaticCode(MODULE_MOTION_SCANNER),
			createModuleByStaticCode(MODULE_GUN),
		},
	})

	startRoom.facilitiesHere = append(startRoom.facilitiesHere,
		&roomFacility{
			code:        FACILITY_GENERATOR,
			number:      1,
			hp:          100,
			destroyable: false,
		},
		&roomFacility{
			code:        FACILITY_INTERFACE,
			number:      1,
			hp:          100,
			destroyable: false,
		},
	)

	for i := 0; i < 8; i++ {
		_, x, y := lvl.getRandomRoomExceptForRange(0, 0, 1, 1)
		lvl.actors = append(lvl.actors, &actor{
			staticId:           ACTOR_MUTANT,
			hp:                 1,
			x:                  x,
			y:                  y,
			isPlayerControlled: false,
		})
	}

	_, x, y := lvl.getRandomRoomExceptForRange(0, 0, 0, 0)
	lvl.rooms[x][y].facilitiesHere = append(lvl.rooms[x][y].facilitiesHere,
		&roomFacility{
			code:        FACILITY_INTERFACE,
			number:      1,
			hp:          100,
			destroyable: false,
		},
	)

	_, x, y = lvl.getRandomRoomExceptForRange(0, 0, 0, 0)
	lvl.rooms[x][y].facilitiesHere = append(lvl.rooms[x][y].facilitiesHere,
		&roomFacility{
			code:        FACILITY_GENERATOR,
			number:      1,
			hp:          100,
			destroyable: false,
		},
	)

	_, x, y = lvl.getRandomRoomExceptForRange(0, 0, 0, 0)
	lvl.rooms[x][y].facilitiesHere = append(lvl.rooms[x][y].facilitiesHere,
		&roomFacility{
			code:        FACILITY_TURRET,
			number:      1,
			hp:          100,
			destroyable: false,
		},
	)

	lvl.currLog = []string{"a", "b", "c"}

	return lvl
}

func (l *level) getRandomRoomInRange(fx, fy, tx, ty int) (*room, int, int) {
	for try := 0; try < 25; try++ {
		x := rnd.RandInRange(fx, tx)
		y := rnd.RandInRange(fy, ty)
		if l.rooms[x][y] != nil {
			return l.rooms[x][y], x, y
		}
	}
	panic("GetRandomRoom failed!")
}

func (l *level) getRandomRoomExceptForRange(fx, fy, tx, ty int) (*room, int, int) {
	for try := 0; try < 25; try++ {
		x := rnd.Rand(len(l.rooms))
		y := rnd.Rand(len(l.rooms[0]))
		if fx <= x && x <= tx && fy <= y && y <= ty {
			continue
		}
		if l.rooms[x][y] != nil {
			return l.rooms[x][y], x, y
		}
	}
	panic("GetRandomRoom failed!")
}
