package piscine

func NotDecimal(dec string) string {
    if dec == "" {
        return "\n"
    }
    
    // Check if has dot
    hasDot := false
    for _, char := range dec {
        if char == '.' {
            hasDot = true
            break
        }
    }
    
    // No dot? Return as is
    if !hasDot {
        return dec + "\n"
    }
    
    // Check if valid number
    for i, char := range dec {
        if char != '.' && char != '-' && (char < '0' || char > '9') {
            return dec + "\n"
        }
        if char == '-' && i != 0 {
            return dec + "\n"
        }
    }
    
    // Remove dot
    result := ""
    for _, char := range dec {
        if char != '.' {
            result += string(char)
        }
    }
    
    // Handle negative
    hasNegative := false
    if result[0] == '-' {
        hasNegative = true
        result = result[1:]
    }
    
    // Remove zeros
    for len(result) > 1 && result[0] == '0' {
        result = result[1:]
    }
    
    // Add negative back
    if hasNegative {
        result = "-" + result
    }
    
    return result + "\n"
}
