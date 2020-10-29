func maxSlidingWindow(nums []int, k int) []int {

    if k == 0 || len(nums) == 0 {
        return nums
    }
    res := make([]int, len(nums)-k+1, len(nums)-k+1)
    maxQue := make([]int,0)

    max := nums[0]
    maxQue = []int{max}
    for i:=1;i<k;i++ {
        if nums[i] > max {
            max = nums[i]
            maxQue = []int{max}
        }else {
            if nums[i] > maxQue[len(maxQue) - 1] {
                for j:=len(maxQue)-1;j>=0;j-- {
                    if maxQue[j] >= nums[i] {
                        maxQue = maxQue[:j+1]
                        break
                    } 
                }
            }
            maxQue = append(maxQue,nums[i])
        }
    }

    res[0] = maxQue[0]
    for i:=k;i<len(nums);i++ {
        if nums[i] > maxQue[0] {
            maxQue = []int{nums[i]}
        }else {
            if nums[i] > maxQue[len(maxQue) - 1] {
                for j:=len(maxQue)-1;j>=0;j-- {
                    if maxQue[j] >= nums[i] {
                        maxQue = maxQue[:j+1]
                        break
                    } 
                }
            }

            maxQue = append(maxQue,nums[i])
        }

        if nums[i-k] == maxQue[0] {
            maxQue = maxQue[1:]
        }
        res[i-k+1] = maxQue[0]
    }
    
    return res
}