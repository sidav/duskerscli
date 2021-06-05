package main

func (l *level) showNetworkTerminal() {
	lines := []string{
		"  Nostromo-OS 3.7182-2283",
		"Username: root",
		"Password: ******",
		"Secure connection established. ",
		"Terminal loaded. Available commands:",
		"It doesn't work yet.",
		"lol",
	}
	rend.showModalTerminal(lines)
}
