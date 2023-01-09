func maximumCount(nums []int) int {
    p:=0
    for i:=0;i<len(nums);i++{
        if nums[i]<0 {
            p++
        } else if nums[i]>0{
            if p>len(nums)-i {
                return p
            } else {
                return len(nums)-i
            }
        }
    }
    return p
}