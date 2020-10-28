func isStraight(nums []int) bool {
    sort.Ints(nums)

    min := 14
    max := -1
    for _,v := range nums {
        if v != 0 {

            if v == max {
                return false
            }

            if v > max {
                max = v
            }

            if v < min  {
                min = v
            }
        }
    }

    return max - min < 5
}