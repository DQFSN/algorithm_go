func singleNumber(nums []int) int {
    dict := make(map[int]int)

    for _,v := range nums {
        dict[v]++
    }
    for _,v := range nums {
        if dict[v] == 1 {
            return v
        }
    }
    
    return -1


}
// 另一种位运算或%3