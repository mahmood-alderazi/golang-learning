package main 

import "fmt"


func main (){


	arry:= []int{4,2,6,1,3,8,7,5}


	pivot:= len(arry)-1 
	
	 i:= -1 

	for j:= 0 ; j < len(arry)-1;j++ {
		if arry[j] < arry[pivot]{
			i++
			arry[i],arry[j] = arry[j],arry[i]
		}
	}
	arry[i+1],arry[pivot] = arry[pivot],arry[i+1]

	fmt.Println(arry)
}
