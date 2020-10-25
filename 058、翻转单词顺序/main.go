func reverseWords(s string) string {
    strArr := make([]string,0,5)

    start := 0
    s = strings.Trim(s, " ")
    for i := 0; i < len(s); i++  {

        if s[i] != ' ' {
            continue
        }else {
            strArr = append(strArr, s[start:i])

            for i < len(s) && s[i] == ' ' {
                i++
            }
            start = i
        }
    }
    strArr = append(strArr, s[start:])
    for s,e := 0,len(strArr)-1;s < e;{
        strArr[s],strArr[e] = strArr[e],strArr[s]
        s++
        e--
    }

    return strings.Join(strArr," ")
}