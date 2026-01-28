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

	slice := []string{}   
	word := ""           

	
	for _, char := range input {
		if char == ' ' {
			if word != "" {
				slice = append(slice, word)
				word = ""
			}
		} else {
			word += string(char)
		}
	}

	// append the last word 
	if word != "" {
		slice = append(slice, word)
	}

	
	for i, w := range slice {
		for _, char := range  w{
			z01.PrintRune(char)
		}
		if i < len(slice)-1 {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}
