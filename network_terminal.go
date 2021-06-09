package main

func (l *level) showNetworkTerminal(termX, termY int) {
	input := ""
	secLevel := l.rooms[termX][termY].securityNumber
	for input != "exit" {
		rend.render(l)
		lines := []string{
			"  Nostromo-OS 3.7182-2283",
			"Username: root",
			"Password: ******",
			"Secure connection established. ",
			"Terminal loaded. Available commands:",
			// TODO: gather available commands
			"survey",
			"unlock",
		}
		input = rend.showModalTerminal(lines)
		l.parseTerminalInput(input, secLevel)
	}
}

func (l *level) parseTerminalInput(input string, secLevel int) {
	switch input {
	case "survey":
		for x := range l.rooms {
			for y := range l.rooms[x] {
				if l.rooms[x][y] != nil && l.rooms[x][y].securityNumber == secLevel {
					l.rooms[x][y].isExplored = true
				}
			}
		}
	}
}
