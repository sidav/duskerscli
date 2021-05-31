package main

import (
	cw "duskerscli/console_wrapper"
	"strconv"
)

type renderer struct {
	roomSizeX, roomSizeY   int // WITH walls!
	logPosition, logHeight int
	drawCoords             bool
	inputLinePosition      int
}

func initRenderer() *renderer {
	cw.Init_console()
	r := &renderer{
		roomSizeX:  6,
		roomSizeY:  4,
		logHeight:  3,
		drawCoords: true,
	}
	r.logPosition = 4*r.roomSizeY + 1
	r.inputLinePosition = r.logPosition+1
	return r
}

func (r *renderer) render(l *level) {
	cw.Clear_console()
	r.renderLevelOutline(l)
	r.renderLevel(l)
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
				chr = rune((strconv.Itoa((y - r.roomSizeY/2) / r.roomSizeY))[0])
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
		} else {
			cw.SetFgColor(cw.RED)
			cw.PutChar(a.getStaticData().char, upx+roomInnerSizeX-enemiesActors-1, upy+1)
		}
	}
}

func (r *renderer) renderLog(l *level) {
	cw.SetFgColor(cw.WHITE)
	cw.PutString(l.currLog, 0, r.logPosition)
}

func (r *renderer) readPlayerInput() string {
	currLine := ""
	key := ""
	for key != "ENTER" {
		currLine, key = cw.ReadTextInputAndKeyPress("> ", currLine, 0, r.inputLinePosition)
	}
	return currLine
}
