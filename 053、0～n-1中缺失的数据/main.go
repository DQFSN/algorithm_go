func missingNumber(nums []int) int {
    // 等差数列求和公式
    left := 0
    right := len(nums)-1

    for left < right {
        mid := (left + right) / 2
        if nums[mid] == mid {
            left = mid + 1
        }else {
            right = mid
        }
    }

    if left == len(nums)-1 && nums[left] == left {
        return len(nums)
    }else {
        return left
    }
}


// func missingNumber(nums []int) int {
//     // 等差数列求和公式
//     sum := 0
//     for _,v := range nums {
//         sum += v
//     }

//     res := (len(nums)+1)*(nums[0] + nums[len(nums)-1])/2 - sum
//     in := false

//     for _,v := range nums {
//         if v == res {
//             in = true
//         }
//     }

//     if in {
//         if nums[len(nums)-1] == len(nums)-1 {
//             return nums[len(nums)-1] + 1
//         }
//         return nums[0] - 1
//     } else {
//         return res
//     }
// }