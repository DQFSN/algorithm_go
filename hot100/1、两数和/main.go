// func twoSum1(nums []int, target int) []int {
    
//     // 数组有序
//     for b,e := 0,len(nums)-1; b < e; {
//         sum := nums[b] + nums[e]
//         if sum == target {
//             return []int{b,e}
//         }else if sum > target {
//             e--
//         }else {
//             b++
//         }
//     }

//     return []int{}
// }

// func twoSum(nums []int, target int) []int {
//     dict := make(map[int]int)
//     for i, v := range nums {
//         dict[v] = i
//     }
//     for i,v := range nums {
//         if r,ok := dict[target - v];ok && i != r  {
//             return []int{r,i}
//         }
//     }
//     return []int{}
// }

func twoSum(nums []int, target int) []int {
    
    dict := make(map[int]int)

    for i, v := range nums {
        dict[v] = i
    }

    for i,v := range nums {
        if r,ok := dict[target - v];ok && i != r  {
            return []int{r,i}
        }
    }

    return []int{}
}