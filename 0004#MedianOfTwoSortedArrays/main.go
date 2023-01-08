func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    p1,p2:=0,0
    merged:=[]int{}
    for p1<len(nums1) && p2<len(nums2) {
        if nums1[p1] < nums2[p2] {
            merged=append(merged,nums1[p1])
            p1++
        } else {
            merged=append(merged,nums2[p2])
            p2++
        }
    }
    for p1=p1;p1<len(nums1);p1++ {merged=append(merged,nums1[p1])}
    for p2=p2;p2<len(nums2);p2++ {merged=append(merged,nums2[p2])}

    if len(merged)%2==0 {return float64(merged[len(merged)/2]+merged[(len(merged)-1)/2])/2} else {return float64(merged[len(merged)/2])}
}