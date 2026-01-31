package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		z01.PrintRune('\n')
		return
	}
	
	input := args[0]
	temp := []string{}
	word := ""
	
	for _, char := range input {
		if char == ' ' {
			if word != "" {
				temp = append(temp, word)
				word = ""
			}
		} else {
			word += string(char)
		}
	}
	
	if word != "" {
		temp = append(temp, word)
	}
	
	if len(temp) == 0 {
		z01.PrintRune('\n')
		return
	}
	
	first := temp[0]  // ← FIXED!
	
	for i := 1; i < len(temp); i++ {  // ← FIXED!
		for _, char := range temp[i] {
			z01.PrintRune(char)
		}
		z01.PrintRune(' ')
	}
	
	for _, char := range first {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
