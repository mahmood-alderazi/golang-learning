package kata


func ReverseList(lst []int) []int {
  
  for i,j:= 0,len(lst)-1 ; i < j ; i,j = i+1,j-1{
    lst[i],lst[j] = lst[j],lst[i]
  }
  
  
  return lst
}
