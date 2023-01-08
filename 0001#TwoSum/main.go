import "fmt"

func twoSum(nums []int, target int) []int {
    // i cannot sort it because then ill mix the initial indices lmao
    // sort.Ints(nums)
    for i:=0;i<len(nums)-1;i++{
        for j:=i+1;j<len(nums);j++{
            if nums[i]+nums[j] == target {return []int{i,j}}
            // break condition doesnt work on unsorted array
            // if nums[i]+nums[j] > target {break}
        }
    }
    return []int{0,0}
}