func nthUglyNumber(n int) int {

    ugly := make([]int,1,n)
    ugly[0] = 1
    two := 0
    three := 0
    five := 0

    for len(ugly) < n {
        twoNum := ugly[two]*2
        threeNum := ugly[three]*3
        fiveNum := ugly[five]*5

        min := findMin(twoNum,threeNum,fiveNum)

        if twoNum == min {
            two++
        }
        if threeNum == min {
            three++
        }
        if fiveNum == min {
            five++
        }
        ugly = append(ugly,min)
    }

    // fmt.Println(ugly)
    return ugly[len(ugly)-1]
}

func findMin(t,th,f int) int {

    if t < th {
        if t < f {
            return t
        }else {
            return f
        }
    }else {
        if th < f {
            return th
        }else {
            return f
        }
    }
}