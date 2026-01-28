package piscine

func RevConcatAlternate(slice1, slice2 []int) []int {
	result := []int{}
	
	len1 := len(slice1)
	len2 := len(slice2)
	maxLen := len1

	if len2 > maxLen {
		maxLen = len2
	}
	
	for i := maxLen; i >= 0; i-- {
		if i < len1 {
			result = append(result, slice1[i])
		}
		if i < len2 {
			result = append(result, slice2[i])
		}
	}
	
	return result
}
