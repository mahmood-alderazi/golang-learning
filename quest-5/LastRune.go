package piscine 

func LastRune(s string) rune {

  runes := []rune(s)

  for i:= len(s)-1; i >= 0 ;i--{
    return runes[i]
  }
  return 0 
}
