package kata

func PositiveSum(numbers []int) int {
  
    sum:= 0 
    for _, num := range numbers{
      
      if num < 0 {
        continue 
      }
      
      sum = sum+ num 
    }
  return sum 
} 
