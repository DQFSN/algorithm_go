func constructArr(a []int) []int {

    if len(a) == 1 {
        return []int{1}
    }
    if len(a) == 0 {
        return nil
    }

    front := make([]int,len(a), len(a))

    front[0] = a[0]
    for i := 1;i < len(a);i++ {
        front[i] = front[i - 1] * a[i]
    }
    back := make([]int,len(a), len(a))

    back[len(a) - 1] = a[len(a) - 1]
    for i:=len(a)-2;i>= 0;i-- {
        back[i] = back[i + 1] * a[i]
    }

    result := make([]int,len(a), len(a))
    result[0] = back[1]
    result[len(a) - 1] = front[len(a) - 2]
    for i:=1;i < len(a)-1;i++ {
        result[i] = front[i-1]*back[i+1]
    }
    return result

}