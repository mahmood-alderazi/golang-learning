package piscine

func FifthAndSkip(str string) string {
    if str == "" {
        return "\n"
    }
    
    // Remove spaces - use runes for multi-byte
    text := []rune{}
    for _, char := range str {
        if char != ' ' {
            text = append(text, char)
        }
    }
    
    // Less than 5? Invalid
    if len(text) < 5 {
        return "Invalid Input\n"
    }
    
    // Take 5, skip 1, repeat
    result := ""
    i := 0
    
    for i < len(text) {
        // Take 5
        count := 0
        for count < 5 && i < len(text) {
            result += string(text[i])
            i++
            count++
        }
        
        // Skip 1 if exists and add space
        if i < len(text) {
            i++
            result += " "
        }
    }
    
    return result + "\n"
}
