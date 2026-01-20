package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if !isValidBase(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	if nbr < 0 {
		z01.PrintRune('-')
		printNegativeInBase(nbr, base)
	} else {
		printInBase(nbr, base)
	}
}

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}

	for i := 0; i < len(base); i++ {
		if base[i] == '+' || base[i] == '-' {
			return false
		}
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				return false
			}
		}
	}

	return true
}

func printInBase(nbr int, base string) {
	baseLen := len(base)

	if nbr >= baseLen {
		printInBase(nbr/baseLen, base)
	}

	z01.PrintRune(rune(base[nbr%baseLen]))
}

func printNegativeInBase(nbr int, base string) {
	baseLen := len(base)

	if nbr <= -baseLen {
		printNegativeInBase(nbr/baseLen, base)
	}

	remainder := -(nbr % baseLen)
	z01.PrintRune(rune(base[remainder]))
}
