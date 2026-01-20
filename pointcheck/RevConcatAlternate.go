package piscine

func RevConcatAlternate(slice1, slice2 []int) []int {
    result := []int{}  // This instead of var result []int
    
    swapped := false
    if len(slice2) > len(slice1) {
        slice1, slice2 = slice2, slice1
        swapped = true
    }
    
    // Add extra elements from slice1 (reverse)
    for i := len(slice1) - 1; i >= len(slice2); i-- {
        result = append(result, slice1[i])
    }
    
    // Alternate the rest (reverse)
    for i := len(slice2) - 1; i >= 0; i-- {
        if swapped  == true  {
            result = append(result, slice2[i])
            result = append(result, slice1[i])
        } else {
            result = append(result, slice1[i])
            result = append(result, slice2[i])
        }
    }
    
    return result
}
