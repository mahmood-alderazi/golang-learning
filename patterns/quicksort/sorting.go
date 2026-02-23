package piscine

func Sorting(arry []int, low int, high int) int {

	pivot := arry[high]

	i := low - 1

	for j := low; j < high; j++ {
		if arry[j] < pivot {
			i++
			arry[i], arry[j] = arry[j], arry[i]
		}
	}

	//final swap

	arry[i+1], arry[high] = arry[high], arry[i+1]

	return i + 1
}

func Quicksort(arry []int, low int, high int) {

	if low >= high {
		return
	}

	pivotindex := Sorting(arry, low, high)

	Quicksort(arry, low, pivotindex-1)
	Quicksort(arry, pivotindex+1, high)

}
