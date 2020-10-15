func replaceSpace(s string) string {
    
    spaceCount := 0
    for _, r := range s {
        if r == ' ' {
            spaceCount++
        }
    }

    newLen := len(s)+spaceCount*2
    ret := make([]byte, newLen, newLen)

    for len := len(s); len != 0; {
        if s[len - 1] == ' ' {
            ret[newLen - 1] = '0'
            newLen--
            ret[newLen - 1] = '2'
            newLen--
            ret[newLen - 1] = '%'
            newLen--
            len--
        }else {
            ret[newLen - 1] = s[len - 1]
            newLen--
            len--
        }
    }

    return string(ret)
}