package ui

import (
	"bufio"
	"os"
)

type Terminal struct {
	reader *bufio.Reader
}

func GetTerminal() *Terminal {
	return &Terminal{
		reader: bufio.NewReader(os.Stdin),
	}
}
