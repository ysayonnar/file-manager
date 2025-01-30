package utils

import (
	"bufio"
	"fmt"
	"os"
)

func IsNumber(b byte) bool {
	s := string(b)
	if s == "0" || s == "1" || s == "2" || s == "3" || s == "4" || s == "5" || s == "6" || s == "7" || s == "8" || s == "9" {
		return true
	}
	return false
}

func ClearStdin() {
	reader := bufio.NewReader(os.Stdin)
	for reader.Buffered() > 0 {
		reader.ReadByte()
	}
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}
