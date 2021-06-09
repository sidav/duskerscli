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
	// place security levels
	secLevels := createFloodFilledIntegerMap(LEVELSIZE, LEVELSIZE, 2)
	// place rooms
	lvl.rooms = make([][]*room, LEVELSIZE)
	for i := range lvl.rooms {
		lvl.rooms[i] = make([]*room, LEVELSIZE)
		for j := range lvl.rooms[i] {
			if !rnd.OneChanceFrom(8) {
				lvl.rooms[i][j] = &room{
					name:           "",
					isExplored:     false,
					isSeenRightNow: false,
					conns:          [2]*connection{},
					securityNumber: secLevels[i][j],
				}
			}
		}
	}
	// place connections
	for x := range lvl.rooms {
		for y := range lvl.rooms[x] {
			if lvl.rooms[x][y] != nil {
				if x < LEVELSIZE-1 && lvl.rooms[x+1][y] != nil && !rnd.OneChanceFrom(5) {
					// to right
					lvl.rooms[x][y].conns[0] = &connection{
						rcx:       1,
						rcy:       0,
						isDoor:    true,
						lockLevel: 0,
						isOpened:  rnd.OneChanceFrom(2) && secLevels[x][y] == secLevels[x+1][y],
						isBroken:  false,
						isLocked:  secLevels[x][y] != secLevels[x+1][y],
					}
				}
				if y < LEVELSIZE-1 && lvl.rooms[x][y+1] != nil && !rnd.OneChanceFrom(5) {
					// to down
					lvl.rooms[x][y].conns[1] = &connection{
						rcx:       0,
						rcy:       1,
						isDoor:    true,
						lockLevel: 0,
						isOpened:  rnd.OneChanceFrom(2) && secLevels[x][y] == secLevels[x][y+1],
						isBroken:  false,
						isLocked:  secLevels[x][y] != secLevels[x][y+1],
					}
				}
			}
		}
	}

	// select start room
	_, srx, sry := lvl.getRandomRoomInRange(1, 0, 0, LEVELSIZE-1, LEVELSIZE-1)

	lvl.actors = append(lvl.actors, &actor{
		name:               "Alpha",
		staticId:           ACTOR_DRONE,
		hp:                 5,
		x:                  srx,
		y:                  sry,
		isPlayerControlled: true,
		modules: []*module{
			createModuleByStaticCode(MODULE_BATTERY),
			createModuleByStaticCode(MODULE_EMERGENCY_GENERATOR),
			createModuleByStaticCode(MODULE_SURVEYOR),
		},
	})
	lvl.actors = append(lvl.actors, &actor{
		name:               "Bravo",
		staticId:           ACTOR_DRONE,
		hp:                 5,
		x:                  srx,
		y:                  sry,
		isPlayerControlled: true,
		modules: []*module{
			createModuleByStaticCode(MODULE_BATTERY),
			createModuleByStaticCode(MODULE_BATTERY),
			createModuleByStaticCode(MODULE_NETWORKING),
			createModuleByStaticCode(MODULE_MOTION_SCANNER),
			createModuleByStaticCode(MODULE_GUN),
		},
	})

	lvl.actors = append(lvl.actors,
		&actor{
			x:  srx,
			y:  sry,
			hp: 100,
			asFacility: &facility{
				code:             FACILITY_GENERATOR,
				associatedNumber: 1,
				destroyable:      false,
			},
		},
		&actor{
			x:  srx,
			y:  sry,
			hp: 100,
			asFacility: &facility{
				code:             FACILITY_INTERFACE,
				associatedNumber: 1,
				destroyable:      false,
			},
		},
	)

	for i := 0; i < 8; i++ {
		_, x, y := lvl.getRandomRoomInRange(-1, 0, 0, LEVELSIZE-1, LEVELSIZE-1)
		lvl.actors = append(lvl.actors, &actor{
			staticId:           ACTOR_MUTANT,
			hp:                 1,
			x:                  x,
			y:                  y,
			isPlayerControlled: false,
		})
	}

	_, x, y := lvl.getRandomRoomInRange(1, 0, 0, LEVELSIZE-1, LEVELSIZE-1)
	lvl.actors = append(lvl.actors,
		&actor{
			x:  x,
			y:  y,
			hp: 100,
			asFacility: &facility{
				code:             FACILITY_GENERATOR,
				associatedNumber: 1,
				destroyable:      false,
			},
		},
	)

	_, x, y = lvl.getRandomRoomInRange(1, 0, 0, LEVELSIZE-1, LEVELSIZE-1)
	lvl.actors = append(lvl.actors,
		&actor{
			x:  x,
			y:  y,
			hp: 100,
			asFacility: &facility{
				code:             FACILITY_INTERFACE,
				associatedNumber: 1,
				destroyable:      false,
			},
		},
	)

	_, x, y = lvl.getRandomRoomInRange(1, 0, 0, LEVELSIZE-1, LEVELSIZE-1)
	lvl.actors = append(lvl.actors,
		&actor{
			x:  x,
			y:  y,
			hp: 100,
			asFacility: &facility{
				code:             FACILITY_TURRET,
				associatedNumber: 1,
				destroyable:      false,
			},
		},
	)

	_, x, y = lvl.getRandomRoomInRange(2, 0, 0, LEVELSIZE-1, LEVELSIZE-1)
	lvl.actors = append(lvl.actors,
		&actor{
			x:  x,
			y:  y,
			hp: 100,
			asFacility: &facility{
				code:             FACILITY_GENERATOR,
				associatedNumber: 2,
				destroyable:      false,
			},
		},
	)

	_, x, y = lvl.getRandomRoomInRange(2, 0, 0, LEVELSIZE-1, LEVELSIZE-1)
	lvl.actors = append(lvl.actors,
		&actor{
			x:  x,
			y:  y,
			hp: 100,
			asFacility: &facility{
				code:             FACILITY_INTERFACE,
				associatedNumber: 2,
				destroyable:      false,
			},
		},
	)

	_, x, y = lvl.getRandomRoomInRange(2, 0, 0, LEVELSIZE-1, LEVELSIZE-1)
	lvl.actors = append(lvl.actors,
		&actor{
			x:  x,
			y:  y,
			hp: 100,
			asFacility: &facility{
				code:             FACILITY_TURRET,
				associatedNumber: 2,
				destroyable:      false,
			},
		},
	)

	lvl.currLog = []string{"a", "b", "c"}

	return lvl
}

func (l *level) getRandomRoomInRange(desiredSecurityNumber, fx, fy, tx, ty int) (*room, int, int) {
	var coords [][2]int
	for x := fx; x <= tx; x++ {
		for y := fy; y <= ty; y++ {
			if l.rooms[x][y] != nil {
				if desiredSecurityNumber == -1 || l.rooms[x][y].securityNumber == desiredSecurityNumber {
					coords = append(coords, [2]int{x, y})
				}
			}
		}
	}
	if len(coords) > 0 {
		ind := rnd.Rand(len(coords))
		x, y := coords[ind][0], coords[ind][1]
		return l.rooms[x][y], x, y
	}
	panic("GetRandomRoom failed!")
}

func (l *level) getRandomRoomExceptForRange(desiredSecurityNumber, fx, fy, tx, ty int) (*room, int, int) {
	var coords [][2]int
	for x := range l.rooms {
		for y := range l.rooms[x] {
			if fx <= x && x <= tx && fy <= y && y <= ty {
				continue
			}
			if l.rooms[x][y] != nil {
				if desiredSecurityNumber == -1 || l.rooms[x][y].securityNumber == desiredSecurityNumber {
					coords = append(coords, [2]int{x, y})
				}
			}
		}
	}
	if len(coords) > 0 {
		ind := rnd.Rand(len(coords))
		x, y := coords[ind][0], coords[ind][1]
		return l.rooms[x][y], x, y
	}
	panic("GetRandomRoomExceptInRange failed!")
}

func (l *level) randomFloodFillSecurities(totalNum int) {
}
