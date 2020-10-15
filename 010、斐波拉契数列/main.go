func fib(n int) int {
    n1, n2 := 0 , 1

    i := 1
    for i <= n {
        n1, n2 = n2, (n1 + n2) % 1000000007 
        i++
    }

    return n1
}