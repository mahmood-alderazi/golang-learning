package piscine

func AtoiBase(s string, base string) int {
	if len(base) < 2 {
		return 0
	}

	for i := 0; i < len(base); i++ {
		if base[i] == '+' || base[i] == '-' {
			return 0
		}
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				return 0
			}
		}
	}

	result := 0

	for _, char := range s {
		found := false
		for i, baseChar := range base {
			if char == baseChar {
				result = result*len(base) + i
				found = true
				break
			}
		}
		if !found {
			return 0
		}
	}

	return result
}
