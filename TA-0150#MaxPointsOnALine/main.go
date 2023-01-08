import "fmt"

func findMax(a []float64) int {
    sort.Float64s(a)
    var max int = 1
    var count int = 1
    var temp float64 = a[0]
    for i:=1;i<len(a);i++{
        if a[i]==temp {
            count++
            if (count > max) {max=count}
        } else {
            count = 1
            temp = a[i]
        }
    }
    return max+1
}

func maxPoints(points [][]int) int {
    if len(points)==1 {return 1}
    var maxLine int = 0
    // all the points are unique!
    // so i can calculate tan(alpha) for every par of points
    for i,point1 := range points {
        var a []float64
        for j,point2 := range points {
            if i == j {
                continue
            } else if point1[0]==point2[0] {
                a = append(a,20001)
            } else {
                a = append(a,float64((point2[1]-point1[1]))/float64((point2[0]-point1[0])))
            }
        }
        fmt.Println(a)
        var tempMax int = findMax(a)
        if tempMax > maxLine {maxLine=tempMax}
    }
    return maxLine
}