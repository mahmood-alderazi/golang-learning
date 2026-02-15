func twoSum(nums []int, target int) []int {

result:= []int{}
for i:= 0 ; i < len(nums); i++{
    for j:= 0 ; j < i ; j++{
        if nums[i] + nums[j] == target{
           result = append(result,j,i)
        }
    }
}  
    return result 
}
