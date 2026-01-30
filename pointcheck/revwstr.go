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
		if char != ' ' {
			word += string(char)
		} else if word != "" {
			temp = append(temp, word)
			word = ""
		}
	}
	
	if word != "" {
		temp = append(temp, word)
	}
	
	result := ""
	for i := len(temp) - 1; i >= 0; i-- {
		result += string(temp[i])
		if i != 0 {
			result += string(" ")
		}
	}
	
	for _, char := range result {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
