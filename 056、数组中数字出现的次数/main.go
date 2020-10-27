func singleNumbers(nums []int) []int {

    var xor int = nums[0]
    for _, v := range nums[1:] {
        xor ^= v
    }

    index := 0
    for xor & 1 != 1{
        xor >>= 1
        index++
    }

    judge := 1 << index


    var first int
    var second int
    ff := true
    sf := true
    for _, v := range nums {
        // 相与后应该等于judge，而不是1
        if v & judge == judge {

            if ff {
                first = v
                ff = false
            }else {
                first ^= v
            }
        }else {
            if sf {
                second = v
                sf = false
            }else {
                second ^= v
            }
        }
    }

    return []int{first,second}
}
