package kata

func CountSheeps(numbers []bool) int {
  
  count:= 0 
  
  for _, sheeps:= range numbers{
      if sheeps == true{
        count++ 
      }
  }
  return count 
}
