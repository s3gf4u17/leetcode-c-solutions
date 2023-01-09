func maximumCount(nums []int) int {
    neg,pos:=[]int{},[]int{}
    for _,val:=range nums {
        if val>0{
            pos=append(pos,val)
        } else if val<0 {
            neg=append(neg,val)
        }
    }
    if len(pos)>len(neg) {
        return len(pos)
    }
    return len(neg)
}