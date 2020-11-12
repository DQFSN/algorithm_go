func lengthOfLongestSubstring(s string) int {
    show := make([]int,1024,1024)

    maxLen := 0
    for i,j:=0,0;j<len(s);{
        index := s[j]
        if show[index] == 0 {
            show[index] = 1
            j++
            if j-i > maxLen {
                maxLen = j-i
            }
        }else {

            for s[i] != s[j] {
                show[s[i]] = 0
                i++
            }
            i++
            show[index] = 0
        }
    }

    return maxLen
}