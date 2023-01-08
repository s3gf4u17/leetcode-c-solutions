func lengthOfLongestSubstring(s string) int {
    if len(s)==0 {return 0}
    p1,p2,max:=0,0,0
    used:=make(map[byte]int)

    for p1<len(s) && p2<len(s) {
        if _,ok:=used[s[p2]];ok {
            delete(used,s[p1])
            p1++
        } else {
            used[s[p2]]=1
            p2++
            if p2-p1>max {max=p2-p1}
        }
    }

    return max
}