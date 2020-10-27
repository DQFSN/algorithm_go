// func findRepeatNumber(nums []int) int {
    
//     temp := make([]int, len(nums), len(nums))
//     if len(nums) == 0 {
//         return -1
//     }
//     temp[0] = -1
//     for _, v := range nums {
//         if temp[v] == v {
//             return v
//         } else {
//             temp[v] = v
//         }
//     }


//     return -1
// }

// 竟然上面的内存消耗更低。。。。

func findRepeatNumber(nums []int) int {
    
    if len(nums) == 0 {
        return -1
    }
    for i, v := range nums {
        if nums[v] == v && i != v {
            return v
        } else {
            nums[v], nums[i] = nums[i], nums[v]
        }
    }


    return -1
}