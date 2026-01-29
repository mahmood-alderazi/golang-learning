package piscine
import (
    "math"
    "strconv"
)
func NotDecimal(dec string) string {
    if dec == "" {
        return "\n"
    }
    
    dotfound := false
    count := 0
    
    for _, char := range dec {
        if char == '.' {
            dotfound = true
        } else if dotfound == true {
            count++
        }
    }
    
    num, err := strconv.ParseFloat(dec, 64)
    if err == nil {
        result := int(num * math.Pow(10, float64(count)))
        return strconv.Itoa(result) + "\n"
    }
    
    return dec + "\n"
}
