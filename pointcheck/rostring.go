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
	
	if args[0] == "" {
		z01.PrintRune('\n')
		return
	}
	
	str := args[0]
	words := []string{}
	word := ""
	
	for _, char := range str {
		if char == ' ' {
			if word != "" {
				words = append(words, word)
				word = ""
			}
		} else {
			word += string(char)
		}
	}
	
	if word != "" {
		words = append(words, word)
	}
	
	first := words[0]
	
	for i := 1; i < len(words); i++ {
		for _, char := range words[i] {
			z01.PrintRune(char)
		}
		z01.PrintRune(' ')
	}
	
	for _, char := range first {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
