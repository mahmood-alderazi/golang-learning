package piscine

func FifthAndSkip(str string) string {
	if str == "" {
		return "\n"
	}
	
	// Remove spaces
	nospace := ""
	for _, char := range str {
		if char != ' ' {
			nospace += string(char)
		}
	}
	
	// Check length
	if len(nospace) < 5 {
		return "Invalid Input\n"
	}
	
	// Take 5, skip 1, repeat
	result := ""
	count := 0
	
	for _, char := range nospace {
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
