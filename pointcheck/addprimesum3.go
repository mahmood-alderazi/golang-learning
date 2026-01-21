package main
import (
    "os"
    "github.com/01-edu/z01"
)
func main() {
    args := os.Args[1:]
    
    // Check 1 arg
    if len(args) != 1 {
        z01.PrintRune('0')
        z01.PrintRune('\n')
        return
    }
    
    // Convert to number
    num := 0
    for _, char := range args[0] {
        if !(char >= '0' || char <= '9') {
            z01.PrintRune('0')
            z01.PrintRune('\n')
            return
        }
        num = num*10 + int(char-'0')
    }
    
    // Sum all primes
        sum := 0
        for i := 2; i <= num; i++ {
            prime := true
            for j := 2; j < i; j++ {
                if i%j == 0 {
                    prime = false
                }
            }
            if prime  == true {
                sum += i
            }
        }
    
    // Print sum
    if sum == 0 {
        z01.PrintRune('0')
    } else {
        result := ""
        for sum > 0 {
        digit := sum % 10
        result = string(rune('0' + digit)) + result
        sum = sum / 10
    }
        for _, char := range result {
            z01.PrintRune(char)
        }
    }
    z01.PrintRune('\n')
}
