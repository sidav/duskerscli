package main

import (
	cw "duskerscli/console_wrapper"
	"fmt"
	"strconv"
)

type renderer struct {
	roomSizeX, roomSizeY    int // WITH walls!
	logYPosition, logHeight int
	drawCoords              bool
	inputLineYPosition      int
	statusXPosition         int
}

func initRenderer() *renderer {
	cw.Init_console()
	r := &renderer{
		roomSizeX:  6,
		roomSizeY:  4,
		logHeight:  3,
		drawCoords: true,
	}
	return r
}

func (r *renderer) render(l *level) {
	r.statusXPosition = 1 + r.roomSizeX*len(l.rooms)
	r.logYPosition = r.roomSizeY*len(l.rooms[0]) + 1
	r.inputLineYPosition = r.logYPosition + 3
	cw.Clear_console()
	r.renderLevelOutline(l)
	r.renderLevel(l)
	r.renderPlayerStatus(l)
	r.renderLog(l)
	cw.SetFgColor(cw.WHITE)
}

func (r *renderer) renderLevelOutline(l *level) {
	cw.SetColor(cw.BLACK, cw.DARK_GRAY)
	for x := 0; x <= r.roomSizeX*len(l.rooms); x++ {
		for y := 0; y <= r.roomSizeY*len(l.rooms[1]); y++ {
			chr := ' '
			if x%(r.roomSizeX) == r.roomSizeX/2 && y == 0 {
				value := 'A' + (x-r.roomSizeX/2)/r.roomSizeX
				chr = rune(value)
			}
			if y%(r.roomSizeY) == r.roomSizeY/2 && x == 0 {
				chr = rune((strconv.Itoa((y-r.roomSizeY/2)/r.roomSizeY + 1))[0])
			}
			// roomCentY := upy + r.roomSizeY/2 - 1
			cw.PutChar(chr, x, y)
		}
	}
	cw.SetBgColor(cw.BLACK)
}

func (r *renderer) renderLevel(l *level) {
	for x := range l.rooms {
		for y := range l.rooms[x] {
			r.renderRoomAt(l, x, y)
		}
	}
}

func (r *renderer) renderRoomAt(l *level, rx, ry int) {
	r.printRoomLetterCoordsAtRoom(l, rx, ry)
	room := l.rooms[rx][ry]
	if room == nil {
		return
	}
	upx := 1 + rx*r.roomSizeX
	upy := 1 + ry*r.roomSizeY
	roomInnerSizeX := r.roomSizeX - 1
	roomInnerSizeY := r.roomSizeY - 1
	if room.isExplored {
		// render room itself
		cw.SetFgColor(cw.DARK_YELLOW)
		roomFloorChr := '.'
		if !room.isSeenRightNow {
			roomFloorChr = ' '
		}
		for x := 0; x < roomInnerSizeX; x++ {
			for y := 0; y < roomInnerSizeY; y++ {
				cw.PutChar(roomFloorChr, upx+x, upy+y)
			}
		}
		if !room.isSeenRightNow {
			r.printRoomLetterCoordsAtRoom(l, rx, ry)
		}
		// render connections
		roomCentX := upx + r.roomSizeX/2 - 1
		roomCentY := upy + r.roomSizeY/2 - 1
		cw.SetFgColor(cw.DARK_MAGENTA)
		for _, c := range room.conns {
			r.renderConnectionForRoomCenterCoords(c, roomCentX, roomCentY, false)
		}
		// also render connections from the other rooms to this one
		r.renderConnectionForRoomCenterCoords(l.getConnBetweenRoomsAtCoords(rx, ry, rx-1, ry), roomCentX, roomCentY, true)
		r.renderConnectionForRoomCenterCoords(l.getConnBetweenRoomsAtCoords(rx, ry, rx, ry-1), roomCentX, roomCentY, true)
	}
	actorsHere := l.getAllActorsAtCoords(rx, ry)
	if room.isSeenRightNow {
		// render actors
		playersActors := 0
		enemiesActors := 0
		facilities := 0
		for _, a := range actorsHere {
			if a.asFacility != nil {
				cw.PutChar(a.getAppearanceChar(), upx+facilities, upy+roomInnerSizeY-1)
				facilities++
			} else if a.isPlayerControlled {
				cw.SetFgColor(cw.DARK_GREEN)
				cw.PutChar(a.getAppearanceChar(), upx+playersActors, upy)
				playersActors++
			} else {
				cw.SetFgColor(cw.RED)
				cw.PutChar(a.getAppearanceChar(), upx+roomInnerSizeX-enemiesActors-1, upy+1)
				enemiesActors++
			}
		}
	} else if room.isUnderMotionScanner {
		cw.SetFgColor(cw.RED)
		cw.PutString(strconv.Itoa(len(actorsHere)), upx+roomInnerSizeX/2, upy+1)
	}
}

func (r *renderer) printRoomLetterCoordsAtRoom(l *level, rx, ry int) {
	room := l.rooms[rx][ry]
	if room == nil || !room.isExplored {
		cw.SetColor(cw.BLACK, cw.DARK_GRAY)
	} else {
		cw.SetColor(cw.DARK_GRAY, cw.BLACK)
	}
	letter := string(rune('A' + rx))
	number := strconv.Itoa(ry+1)
	roomCentX := 1 + rx*r.roomSizeX + r.roomSizeX/2 - 1
	roomCentY := 1 + ry*r.roomSizeY + r.roomSizeY/2 - 1
	if r.roomSizeX % 2 != 0 {
		cw.PutString(letter+number, roomCentX, roomCentY)
	} else {
		cw.PutString(letter+"-"+number, roomCentX-1, roomCentY)
	}
	cw.SetColor(cw.WHITE, cw.BLACK)
}

func (r *renderer) renderConnectionForRoomCenterCoords(c *connection, rx, ry int, reverse bool) {
	if c == nil {
		return
	}
	chr := '+'
	if c.isOpened {
		chr = '\''
	}
	if c.isBroken {
		chr = '\\'
	}
	if reverse {
		cw.PutChar(chr, rx-c.rcx*(r.roomSizeX/2), ry-c.rcy*r.roomSizeY/2)
	} else {
		cw.PutChar(chr, rx+c.rcx*(r.roomSizeX/2), ry+c.rcy*r.roomSizeY/2)
	}
}

func (r *renderer) renderPlayerStatus(l *level) {
	cw.SetFgColor(cw.WHITE)
	cw.PutString(fmt.Sprintf("Turn %d.%d", l.currentTurnNumber/10, l.currentTurnNumber%10),
		r.statusXPosition, 0)
	currY := 1
	for _, a := range l.actors {
		if a.isPlayerControlled {
			strsToPut := make([]string, 0)
			strsToPut = append(strsToPut, fmt.Sprintf("%s: \"%s\"", a.getStaticData().defaultName, a.name))
			//currEnergy, maxEnergy := a.getEnergyCurrAndMax()
			//strsToPut = append(strsToPut, fmt.Sprintf("ENERGY: %d/%d", currEnergy, maxEnergy))
			for i := range strsToPut {
				cw.PutString(strsToPut[i], r.statusXPosition, currY)
				currY++
			}
			for i := range a.modules {
				if a.modules[i].isEnabled {
					cw.SetColor(cw.BLACK, cw.DARK_GREEN)
				}
				cw.PutString(a.modules[i].getNameAndEnabled(), r.statusXPosition, currY)
				currY++
				cw.SetColor(cw.WHITE, cw.BLACK)
			}
			currY++
		}
	}
}

func (r *renderer) renderLog(l *level) {
	cw.SetFgColor(cw.WHITE)
	for i := range l.currLog {
		cw.PutString(l.currLog[i], 0, r.logYPosition+i)
	}
}

func (r *renderer) showModalTerminal(lines []string) string {
	cwid, chei := cw.GetConsoleSize()
	winx := r.statusXPosition
	winy := 0
	winw := cwid - winx - 1
	winh := chei - 1
	for x := winx; x < winx+winw; x++ {
		for y := winy; y < winy+winh; y++ {
			cw.PutChar(' ', x, y)
		}
	}
	cw.SetBgColor(cw.DARK_GREEN)
	for x := winx; x <= winx+winw; x++ {
		cw.PutChar(' ', x, winy)
		cw.PutChar(' ', x, winy+winh)
	}
	for y := winy; y <= winy+winh; y++ {
		cw.PutChar(' ', winx, y)
		cw.PutChar(' ', winx+winw, y)
	}
	cw.SetColor(cw.DARK_GREEN, cw.BLACK)
	for i := range lines {
		cw.PutString(lines[i], winx+1, winy+1+i)
	}
	currLine := ""
	key := ""
	for key != "ENTER" {
		currLine, key = cw.ReadTextInputAndKeyPress("> ", currLine, winx+1, winy+len(lines)+2)
		if key == "CTRL+C" {
			abortGame = true
			return "exit"
		}
	}
	return currLine
}

func (r *renderer) readPlayerInput() string {
	currLine := ""
	key := ""
	for key != "ENTER" {
		currLine, key = cw.ReadTextInputAndKeyPress("> ", currLine, 0, r.inputLineYPosition)
		if key == "CTRL+C" {
			abortGame = true
			return "exit"
		}
	}
	return currLine
}
