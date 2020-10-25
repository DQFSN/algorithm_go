



func translateNum(num int) int {
    nums := make([]int,0,5)

    for num != 0 {
        nums = append(nums, num % 10)
        num /= 10
    }

    for s,e:= 0,len(nums)-1;s<e;{
        nums[s],nums[e] = nums[e],nums[s]
        s++
        e--
    }

    // 去重复
    results := make(map[string]bool)

    dfs(nums,[]byte{},&results)


    // fmt.Println(results)
    // for _,x := range results {
    //     fmt.Printf("%s\t",x)
    // }
    return len(results)
}

func dfs(nums []int, res []byte, results *map[string]bool) {
    if len(nums) == 0 {
        // 还是之前的那个内存坑，如果不拷贝，之后的结果会覆盖掉之前的结果
        // tmp := make([]byte,len(res),len(res))
        // copy(tmp,res)
        // *results = append(*results,tmp)
        
        (*results)[string(res)] = true

        return
    }

    if len(nums) >= 1 {
        res = append(res, byte(96 + nums[0]))
        dfs(nums[1:],res,results)
    }

    if len(nums) >= 2 {
        tmp := nums[0]*10 + nums[1]
        if tmp < 26 {
            res = append(res, byte(96 + tmp))
            dfs(nums[2:],res,results)
        }
    }


}