package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		return
	}

	upper := false
	if args[0] == "--upper" {
		upper = true
		args = args[1:]
	}

	for _, arg := range args {
		num := 0
		for _, c := range arg {
			if c >= '0' && c <= '9' {
				num = num*10 + int(c-'0')
			} else {
				num = 0
				break
			}
		}

		if num >= 1 && num <= 26 {
			if upper {
				z01.PrintRune(rune('A' + num - 1))
			} else {
				z01.PrintRune(rune('a' + num - 1))
			}
		} else {
			z01.PrintRune(' ')
		}
	}

	z01.PrintRune('\n')
}
