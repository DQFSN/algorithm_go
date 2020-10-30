func search(nums []int, target int) int {

    left := 0
    right := len(nums)-1

    for left <= right {
        mid := (left+right) / 2
        fmt.Println(left,right,nums[mid])
        if nums[mid] == target {
            count := 0

            for i := mid;i<=right; i++ {
                if nums[i] == target {
                    count++
                } 
            }
            for i := mid-1;i>=left; i-- {
                if nums[i] == target {
                    count++
                } 
            }
            return count

        }else if nums[mid] < target {
            left = mid + 1
        }else {
            right = mid - 1
        }
    }

    return 0
}