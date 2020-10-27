func minNumber(nums []int) string {

    numsStr := make([]string, 0, len(nums))

    for _,v := range nums {
        str := strconv.Itoa(v)
        numsStr = append(numsStr, str)
    }
    return strings.Join(StrSort(numsStr), "")

}

// 直接根据题意比较
func strCompare(a, b string) bool {

    return a+b < b+a

}

// 不仅低效而且还有解决不了的case:824、8247
func strCompare1(a, b string) bool {

    maxLen := len(a) 
    if len(a) < len(b) {
        maxLen = len(b)
    }

    for i:=0;i<maxLen;{
        ac := byte(' ')
        bc := byte(' ')
        if i >= len(a) {
            ac = a[len(a) - 1]
        }else {
            ac = a[i]
        }

        if i >= len(b) {
            bc = b[len(b) - 1]
        }else {
            bc = b[i]
        }

        if ac == bc {
            i++
            continue
        }else {
            return ac < bc
        }
        
    }

    return false

}

func StrSort(strs []string) (res []string) {
    res = make([]string, len(strs), len(strs))
    slected := make([]int, len(strs), len(strs))

    for j,_ := range strs {
        min := "AA"
        minIndex := 0
        for i, v := range strs {
            if slected[i] == 1 {
                continue
            }
            if strCompare(v,min) {
                min = v
                minIndex = i
            }
        }
        slected[minIndex] = 1 
        res[j] = min
    }

    return
}