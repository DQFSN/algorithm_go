func firstUniqChar(s string) byte {
    // order := make([]rune,0,26)
    dict := make(map[rune]int)

    for _,v := range s {
        // if dict[v] == 0 {
        //     order = append(order,v)
        // }
        dict[v]++
    }

    for _,v := range s {
    // for _,v := range order {
        if dict[v] == 1 {
            return byte(v)
        }
    }

    return ' '
}

// 直接数组作为map更快