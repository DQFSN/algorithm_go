// func numWays(n int) int {

//     if n <= 1 {
//         return 1
//     }

//     return (numWays(n - 1) + numWays(n - 2)) % 1000000007
// }

func numWays(n int) int {

    n0, n1 := 1, 1

    for i := 2; i <= n; i++ {
        n0, n1 = n1, (n0 + n1) % 1000000007
    } 

    return n1
}