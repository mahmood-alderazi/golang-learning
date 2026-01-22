package piscine

import (
    "strconv"
)

func NotDecimal(dec string) string {
    // Empty check
    if dec == "" {
        return "\n"
    }
    
    // Find decimal point
    hasDot := false
    for _, char := range dec {
        if char == '.' {
            hasDot = true
            break
        }
    }
    
    // No decimal point? Return as is
    if !hasDot {
        return dec + "\n"
    }
    
    // Check if only zeros after decimal
    onlyZeros := false 
    afterDot := false
    for _, char := range dec {
        if char == '.' {
            onlyZeros = true 
            afterDot = true
        }
        if afterDot && char != '0' && char != '.' {
            onlyZeros = false
        }
    }
    
    // Only zeros? Return as is
    if onlyZeros {
        return dec + "\n"
    }
    
    // Remove decimal point
    clean := ""
    for _, char := range dec {
        if char != '.' {
            clean += string(char)
        }
    }
    
    // Convert to int
    num, err := strconv.Atoi(clean)
    if err != nil {
        return dec + "\n"
    }
    
    return strconv.Itoa(num) + "\n"
}
