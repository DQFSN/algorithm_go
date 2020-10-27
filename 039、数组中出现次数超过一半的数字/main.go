// 一换一，人头优势
func majorityElement(nums []int) int {
    fighter := 0
    count := 0
    for _, enemy := range nums {
        if count == 0 {
            fighter = enemy
            count++
            continue
        }

        if enemy != fighter {
            count--
        }else {
            count++
        }
    }

    return fighter
}

// 排序中间数

// func majorityElement(nums []int) int {
//     return findKthLarge(nums,0,len(nums) - 1,len(nums)/2)
// }

// func findKthLarge(nums []int, start, end, k int) int {

//     l := start
//     r := end

//     if start >= end {
//         return nums[start]
//     }

//     piovt := nums[end]
//     for start != end {
//         for nums[start] < piovt {
//             start++
//         }

//         for end > start && nums[end] >= piovt {
//             end--
//         }

//         nums[start],nums[end] = nums[end],nums[start]
//     }

//     nums[start],nums[r] = nums[r],nums[start]

//     if k == start {
//         return nums[start]
//     }else if k < start {
//         return findKthLarge(nums,l,start - 1,k)
//     }
    
//     return findKthLarge(nums,start+1,r,k)
    
// }