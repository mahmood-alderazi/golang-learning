package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
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
	
	for i := len(temp) - 1; i >= 0; i-- {
		for _, char := range temp[i] {
			z01.PrintRune(char)
		}
		if i > 0 {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}
