package kata

func ToAlternatingCase(str string) string {
    
      result:= ""
      for _, char := range str{
        if char >= 'A' && char <= 'Z'{
          result+= string(rune(int(char) + 32))
        }else if char >= 'a' && char <= 'z'{
          result+= string(rune(int(char) - 32))
        }else{
          result+= string(char)
        }
      }
    return string(result) 
}
