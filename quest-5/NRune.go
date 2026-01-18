package piscine

func NRune(s string, n int) rune {
	if n <= 0 {
		return 0
	}

	count := 1
	for _, char := range s {
		if n == count {
			return char
		}
		count++
	}
	return 0
}
