func lastRemaining(n int, m int) int {

    res := 0

    for i :=2;i<=n;i++ {
        res = (m + res) % i 
    }

    return res
}


// 递归
// func lastRemaining(n int, m int) int {

//     if n == 1 {
//         return 0
//     }

//     pre := lastRemaining(n - 1, m)
//     // fmt.Println(pre)
//     return (m + pre) % n
// }


// 超时
// func lastRemaining(n int, m int) int {

//     nums := make([]int,n,n)
//     var count int
//     tickOut := make(map[int]bool)
//     var res int
//     for i:= 0;true;i = (i+1)%n {
//         if nums[i] == 0 {
//             count++
//             if count % m == 0 {
//                 nums[i] = 1
//                 res = i
//                 tickOut[i] = true
//                 // fmt.Printf("%d %v \n",res,nums)
//                 count = 0
//             }
//         }
        
//         if len(tickOut) == n {
//             break
//         }

//     }

//     return res
// }