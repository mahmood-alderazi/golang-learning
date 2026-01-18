package piscine

func FirstRune(s string) rune {
	if len(s) == 0 {
		return 0
	}

	for _, char := range s {
		return char
	}
	return 0
}
