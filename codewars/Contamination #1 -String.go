package kata


func Contamination(text, char string) string {
  
  //If the text or the character are empty, return an empty string.
    if text == "" || char == ""{
      return ""
    }
    
   
  result:= ""
  for i:= 0 ; i <len(text);i++{
    result+= string(char)
  }
    return result
}
