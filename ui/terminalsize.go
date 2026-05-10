package ui

import (
	"os"

	"golang.org/x/term"
)

func TerminalSize() (width, height int) {
	w, h, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return 80, 30
	}

	return w, h
}
