func maxSubArray(nums []int) int {
        
    sum := 0
    maxSum := 0
    // 全为负数时返回最大值，题目要求至少有一个元素
    maxNum := nums[0]
    for _,v := range nums {
        if sum + v >= 0 {
            sum += v
            if sum > maxSum {
                maxSum = sum
            }
        }else {
            sum = 0
        }

        if v > maxNum {
            maxNum = v
        }

    }

    if  maxSum <= 0 {
        return maxNum
    }
    return maxSum
}