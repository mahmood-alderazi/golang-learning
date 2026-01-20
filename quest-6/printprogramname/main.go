package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	programPath := os.Args[0]

	lastSeparator := -1
	for i := 0; i < len(programPath); i++ {
		if programPath[i] == '/' {
			lastSeparator = i
		}
	}

	fileName := programPath[lastSeparator+1:]

	for _, char := range fileName {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
