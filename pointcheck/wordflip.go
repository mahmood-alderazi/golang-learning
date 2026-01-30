package piscine

import "strings"

func WordFlip(str string) string {
    if str == "" {
        return "Invalid Output\n"
    }
    
    words := strings.Split(str, " ")
    result := ""
    
    for i := len(words)-1 ; i >= 0; i-- {
        if words[i] != "" {
            result += words[i] + " "
        }
    }
    
    result = strings.TrimSpace(result)
    return result + "\n"
}
