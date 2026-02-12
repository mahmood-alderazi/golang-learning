package kata


func CountBy(x, n int) []int {
  
  result:= []int{}
    
 for i:= x ; i <= n*x ; i+=x{
		result = append(result,i)
	}
    return result 
}
