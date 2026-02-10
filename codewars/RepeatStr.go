package kata

func RepeatStr(repetitions int, value string) string {
  
  if repetitions < 0 {
    return ""
  } 
  
  if value == "" {
    return ""
  }
    
  
   result:= ""
  for i:= 0 ; i < repetitions; i++{
    result+= string(value)
  }
  return result
}
