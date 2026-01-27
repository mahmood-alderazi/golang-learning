package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	arg := os.Args[1:]

	if len(arg) != 2 {
		fmt.Println("Invalid input.")
		return
	}

	s1 := arg[0]
	s2 := arg[1]

	if s1[0] != '[' || s1[len(s1)-1] != ']' {
		fmt.Println("Invalid input.")
		return
	}

	arr := strings.Split(s1[1:len(s1)-1], ", ")


	NumArr := []int{}
	for _, char := range arr {
		temp, err := strconv.Atoi(char)
		if err != nil {
			fmt.Printf("Invalid number: %s\n", char)
			return
		}
		NumArr = append(NumArr, temp)
	}

	target, err := strconv.Atoi(s2)
	if err != nil {
		fmt.Println("Invalid target sum.")
		return
	}
	var indexSum [][]int

	for i := 0; i < len(NumArr)-1; i++ {
		for j := i + 1; j < len(NumArr); j++ {
			if NumArr[i]+NumArr[j] == target {
				indexSum = append(indexSum, []int{i, j})
			}
		}
	}
	if len(indexSum) == 0 {
		fmt.Println("No pairs found.")
		return
	}
	fmt.Printf("Pairs with sum %d: %v\n", target, indexSum)
}
