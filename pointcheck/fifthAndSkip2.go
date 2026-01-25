package piscine

func FifthAndSkip(str string) string {
	if str == "" {
		return "\n"
	}

	// Remove spaces
	clearStr := ""
	for _, char := range str {
		if char != ' ' {
			clearStr += string(char)
		}
	}

	// Check length
	if len(clearStr) < 5 {
		return "Invalid Input\n"
	}

	// Take 5, skip 1, repeat
	result := ""
	count := 0

	for _, char := range clearStr {
		if count == 5 {
			result += " "
			count = 0
		} else {
			result += string(char)
			count++
		}
	}

	return result + "\n"
}
