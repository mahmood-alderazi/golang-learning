func isAnagram(s string, t string) bool {

	if len(s) != len(t){
		return false 
	}

	char:= []rune(t)

	char2 := []rune(s)

	for i:= 0 ; i < len(char); i++{
		for j:= i+1 ; j < len(char);j++ {
			if char[i] > char[j]{
				char[i],char[j] = char[j],char[i]
			}
		}
	}

	for i:= 0 ; i < len(char2); i++{
		for j:= i+1 ; j < len(char2);j++ {
			if char2[i] > char2[j]{
				char2[i],char2[j] = char2[j],char2[i]
			}
		}
	}
	

	result1:= string(char)
	result2:= string(char2)

	if result1 == result2{
		return true 
	}
	return false 
}
