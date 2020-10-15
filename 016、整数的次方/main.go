func myPow(x float64, n int) float64 {

    neg := false
    if n < 0 {
        n = -n
        neg = true
    }
    base := x
    result := 1.0
    for n != 0 {
        // 当前为1是才更新最后结果
        if n & 1 == 1 {
            result = result * base
        }

        base = base * base
        n = n >> 1 
    }

    if neg {
        result = 1.0 / result
    }

    return result

}