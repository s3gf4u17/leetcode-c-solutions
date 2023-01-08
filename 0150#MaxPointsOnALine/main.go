func maxPoints(points [][]int) int {
    var maxLength int = 1
    xs:=make(map[int]int)
    for i:=0;i<len(points);i++{
        as:=make(map[float32]int)
        for j:=i+1;j<len(points);j++{
            a:=float32(points[j][1]-points[i][1])/float32(points[j][0]-points[i][0])
            if val,ok:=as[a];ok{ as[a]=val+1 } else { as[a]=2 }
            if as[a]>maxLength {maxLength=as[a]}
        }
        if val,ok:=xs[points[i][0]];ok{ xs[points[i][0]]=val+1 } else { xs[points[i][0]]=1 }
        if xs[points[i][0]]>maxLength {maxLength=xs[points[i][0]]}
    }
    return maxLength
}