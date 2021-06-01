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
	r.logYPosition = r.logHeight + r.roomSizeY*4
	r.inputLineYPosition = r.logYPosition + 1
	r.statusXPosition = 1 + r.roomSizeX*4
	return r
}

func (r *renderer) render(l *level) {
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
			if x%(r.roomSizeX) == r.roomSizeX/2 {
				value := 'A' + (x-r.roomSizeX/2)/r.roomSizeX
				chr = rune(value)
			}
			if y%(r.roomSizeY) == r.roomSizeY/2 {
				chr = rune((strconv.Itoa((y - r.roomSizeY/2) / r.roomSizeY + 1))[0])
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
	room := l.rooms[rx][ry]
	if room == nil {
		return
	}
	upx := 1 + rx*r.roomSizeX
	upy := 1 + ry*r.roomSizeY
	roomInnerSizeX := r.roomSizeX - 1
	roomInnerSizeY := r.roomSizeY - 1
	// render room itself
	for x := 0; x < roomInnerSizeX; x++ {
		for y := 0; y < roomInnerSizeY; y++ {
			cw.PutChar(' ', upx+x, upy+y)
		}
	}
	// render connections
	roomCentX := upx + r.roomSizeX/2 - 1
	roomCentY := upy + r.roomSizeY/2 - 1
	cw.SetFgColor(cw.DARK_MAGENTA)
	for _, c := range room.conns {
		if c != nil {
			chr := '+'
			if c.isOpened {
				chr = '\''
			}
			if c.isBroken {
				chr = '\\'
			}
			cw.PutChar(chr, roomCentX+c.rcx*(r.roomSizeX/2), roomCentY+c.rcy*r.roomSizeY/2)
		}
	}
	// render actors
	actorsHere := l.getAllActorsAtCoords(rx, ry)
	playersActors := 0
	enemiesActors := 0
	for _, a := range actorsHere {
		if a.isPlayerControlled {
			cw.SetFgColor(cw.DARK_GREEN)
			cw.PutChar(a.getStaticData().char, upx+playersActors, upy)
			playersActors++
		} else {
			cw.SetFgColor(cw.RED)
			cw.PutChar(a.getStaticData().char, upx+roomInnerSizeX-enemiesActors-1, upy+1)
			enemiesActors++
		}
	}
}

func (r *renderer) renderPlayerStatus(l *level) {
	cw.SetFgColor(cw.WHITE)
	cw.PutString(fmt.Sprintf("Turn %d", l.currentTurnNumber), r.statusXPosition, 0)
	currY := 1
	for _, a := range l.actors {
		if a.isPlayerControlled {
			cw.PutString(fmt.Sprintf("%s: %s", a.getStaticData().defaultName, a.name),
				r.statusXPosition, currY)
			currY++
		}
	}
}

func (r *renderer) renderLog(l *level) {
	cw.SetFgColor(cw.WHITE)
	cw.PutString(l.currLog, 0, r.logYPosition)
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
