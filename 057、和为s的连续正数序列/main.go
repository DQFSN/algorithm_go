func findContinuousSequence(target int) [][]int {

    max := target/2 + 1
    arr := make([]int, max+1, max+1)
    for i:=0;i<len(arr);{
        arr[i]=i
        i++
    }

    result := make([][]int,0)

    for front,back := 1,2; front <= back && back <= max;  {
        s := sum(front,back)
        if s == target {
            // tmp := make([]int, back - front+1, back - front + 1)
            // copy(tmp,arr[front:back+1])
            // result = append(result, tmp)
            result = append(result, arr[front:back+1])
            front++
        }else if s < target {
            back++
        }else {
            front++
        }
    }

    return result
}

func sum(s, e int) int {
    // sum := 0
    // for s <= e {
    //     sum += s
    //     s++
    // }

    sum := (e+s)*(e-s+1) / 2

    return sum
}