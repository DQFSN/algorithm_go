func permutation(s string) []string {

    chars := []rune(s)

    // fmt.Printf("%v\n",string(chars))
    result := make([][]rune, 0, 20)

    for i := 0; i < len(chars); i++ {
        
        newResults := make([][]rune, 0)

        for _, r := range result {
            r = append(r, chars[i])

            for j, _ := range r {
                // 记录原始的
                srcR := r
                // fmt.Printf("r---%v\t", string(r))
                // r 为引用类型，需拷贝
                copy := make([]rune, len(r), len(r))
                for i, t := range r {
                    copy[i] = t
                }

                // copy := []rune(r)
                r = copy

                // tmp := r[j]
                // r[j] =  r[len(r) - 1]
                // r[len(r) - 1] = tmp
                r[j], r[len(r) - 1] = r[len(r) - 1], r[j]
                newResults = append(newResults, r)
                r = srcR
            }
        }

        if len(result) == 0 {
            first := []rune{chars[0]}
            result = append(result, first)
        }

        for _, nr := range newResults {
            result = append(result, nr)
            // fmt.Printf("%v\t", string(nr))
        }
        // fmt.Printf("\n")

    }

    resultString := make([]string, 0, len(result)/2)
    mapStr := make(map[string]bool)
    for _, r := range result {
        if len(r) == len(chars) {
            if !mapStr[string(r)] {
                mapStr[string(r)] = true
                resultString = append(resultString, string(r))
            }
        }
    }

    if len(resultString) == 0 {
        return []string{""}
    }

    return resultString
}