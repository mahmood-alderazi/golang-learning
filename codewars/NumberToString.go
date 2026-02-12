package kata


func NumberToString(n int) string {
 
  if n == 0 {
    return "0"
  }
    
  isnegative:= false 
  if n < 0{
     isnegative =  true  
    n = -n 
  }
  
  
  result:= ""
  
  for n > 0 {
    lastnumber:= n%10
    result = string(rune(lastnumber + '0')) + result 
    n = n / 10 
  }
    
  
  if  isnegative == true{
    return "-" + result 
  }else{
    return result 
  }
  return result 
}
