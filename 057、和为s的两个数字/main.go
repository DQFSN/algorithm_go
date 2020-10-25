func twoSum(nums []int, target int) []int {


    for s, e := 0, len(nums)-1;s < e; {

        sum := nums[s] + nums[e]
        if sum == target {
            return []int{nums[s],nums[e]}
        }else if sum < target {
            s++
        }else {
            e--
        }
    } 

    return nil
}