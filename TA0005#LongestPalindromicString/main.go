func isPalindrome(s string) bool {
    if len(s)==1 {return true}
    pL,pR:=0,len(s)-1
    for pL!=pR && pL<pR {
        if s[pL]!=s[pR] {return false}
        pL++
        pR--
    }
    return true
}

func longestPalindrome(s string) string {
    substring:=""
    for i:=0;i<len(s);i++{
        for j:=i;j<len(s);j++{
            if isPalindrome(s[i:j+1]) && j-i+1>len(substring){substring=s[i:j+1]}
        }
    }
    return substring
}