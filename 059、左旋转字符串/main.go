func reverseLeftWords(s string, n int) string {
    byts := []byte(s)
    pre := make([]byte,n,n)
    copy(pre,byts[:n])
    for i:=n;i<len(s);i++{
        byts[i-n] = byts[i]
    }
    for i,v := range pre {
        byts[len(s) - n + i] = v
    }
    
    return string(byts)
}