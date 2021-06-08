package main

func (l *level) createPathfindingCostMap() [][]int {
	w, h := len(l.rooms), len(l.rooms[0])
	costMap := make([][]int, 2*w+1)
	for i := range costMap {
		costMap[i] = make([]int, 2*h+1)
		for j := range costMap[i] {
			costMap[i][j] = -1
		}
	}

	for x := range l.rooms {
		for y := range l.rooms[x] {
			if l.rooms[x][y] != nil {
				costMap[1+x*2][1+y*2] = 1
				for _, c := range l.rooms[x][y].conns {
					if c != nil && c.isOpened {
						costMap[1+x*2+c.rcx][1+y*2+c.rcy] = 0
					}
				}
			}
		}
	}
	return costMap
}

func (l *level) findNextStepInPathFromTo(fx, fy, tx, ty int) (int, int) {
	costmap := l.createPathfindingCostMap()
	path := l.pathfinder.FindPath(&costmap, 1+2*fx, 1+2*fy, 1+2*tx, 1+2*ty)
	return path.GetNextStepVector()
}
