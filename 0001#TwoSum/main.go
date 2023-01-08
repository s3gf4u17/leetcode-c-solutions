func twoSum(nums []int, target int) []int {
    hashtable := make(map[int]int)
    for i,num := range nums {
        if index,ok:=hashtable[num];ok {
            return []int{index,i}
        }
        hashtable[target-num]=i
    }
    return []int{0,0}
}