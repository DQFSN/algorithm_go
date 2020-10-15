func exchange(nums []int) []int {
    low,high := 0, len(nums) - 1

    for low < high {
        for nums[low] % 2 != 0 && low < high{
            low++
        }
        for nums[high] % 2 != 1 && low < high{
            high--
        }

        temp := nums[low]
        nums[low] = nums[high]
        nums[high] = temp

        low++
        high--

    }

    return nums
}