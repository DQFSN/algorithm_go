func lengthOfLongestSubstring(s string) int {

    dict := make(map[rune]bool)
    maxLen := 0

    l := 0
    start := 0
    for i:=0;i<len(s);i++ {
        
        v := rune(s[i])

        if dict[v] == false {
            dict[v] = true
            l++
            continue
        }else {
            if l > maxLen {
                maxLen = l
            }

            for v != rune(s[start]) {
                dict[rune(s[start])] = false
                start++
                l--
            }
            start++

        }
    }

    if l >  maxLen {
        maxLen = l
    }

    return maxLen
}