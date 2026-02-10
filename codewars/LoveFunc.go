package kata

func LoveFunc(flower1, flower2 int) bool {
  
  
  if flower1%2 ==  0 && !(flower2%2  == 0){
    return true   
  }else if !(flower1 % 2 == 0) && flower2%2 == 0  {
    return true 
  }
  
  
  return false  
}
