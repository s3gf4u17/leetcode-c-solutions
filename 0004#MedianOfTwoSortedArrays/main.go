func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    p1,p2:=0,0
    merged:=[]int{}
    for p1<len(nums1) || p2<len(nums2) {
        if p1>=len(nums1) && p2<len(nums2) {
            merged=append(merged,nums2[p2])
            p2++
        } else if p1<len(nums1) && p2>=len(nums2) {
            merged=append(merged,nums1[p1])
            p1++
        } else if p2<len(nums2) && p1<len(nums1) {
            if nums1[p1]>nums2[p2] {
                merged=append(merged,nums2[p2])
                p2++
            } else {
                merged=append(merged,nums1[p1])
                p1++
            }
        } else {
            break
        }
    }
    if (len(nums1)+len(nums2))%2==0 {
        return float64(merged[(len(nums1)+len(nums2))/2]+merged[(len(nums1)+len(nums2))/2-1])/2
    } else {
        return float64(merged[(len(nums1)+len(nums2))/2])
    }
}