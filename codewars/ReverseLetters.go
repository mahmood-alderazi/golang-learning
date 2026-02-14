package kata


func ReverseLetters(s string) string {
  
  if s == ""{
    return ""
  }
    
  
  result:= ""
  for _, char := range s {
    if char >= 'a' && char <= 'z'{
      result+= string(char)
    }else if char >= 'A' && char <= 'Z'{
      result+= string(char)
    }
  }
    
  rev:= ""
  for i:= len(result)-1; i >= 0 ; i--{
    rev+= string(result[i])
  } 
  return rev
}
