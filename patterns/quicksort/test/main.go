package main 

import (
	"fmt"
	"piscine"
)


func main(){

	arry := []int{4, 2, 6, 1, 3, 8, 7, 5}

	piscine.Quicksort(arry,0,len(arry)-1)


	fmt.Print(arry)
}
