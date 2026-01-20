package piscine

func Capitalize(s string) string {
	result := ""
	newWord := true

	for _, char := range s {
		if isAlphaNum(char) {
			if newWord {
				result += toUpper(char)
				newWord = false
			} else {
				result += toLower(char)
			}
		} else {
			result += string(char)
			newWord = true
		}
	}
	return result
}

func isAlphaNum(char rune) bool {
	if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
		return true
	}
	return false
}

func toUpper(char rune) string {
	if char >= 'a' && char <= 'z' {
		return string(char - 32)
	}
	return string(char)
}

func toLower(char rune) string {
	if char >= 'A' && char <= 'Z' {
		return string(char + 32)
	}
	return string(char)
}
