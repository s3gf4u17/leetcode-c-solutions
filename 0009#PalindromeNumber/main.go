func isPalindrome(x int) bool {
    if x < 0 {return false}

    // avoid overlow
    var y int64 = 0
    var z int64 = int64(x)

    for z > 0 {
        y = y * 10 + z % 10
        z /= 10
    }

    return y == int64(x)
}