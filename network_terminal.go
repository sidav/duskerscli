package main

func (l *level) showNetworkTerminal() {
	input := ""
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
		l.parseTerminalInput(input)
	}
}

func (l *level) parseTerminalInput(input string) {
	switch input {
	case "survey":
		for x := range l.rooms {
			for y := range l.rooms[x] {
				l.rooms[x][y].isExplored = true
			}
		}
	}
}
