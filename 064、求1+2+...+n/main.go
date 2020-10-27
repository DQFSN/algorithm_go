package _64_求1_2_____n


func sumNums(n int) int {
	sum := n

	f := func () int {
		sum = n + sumNums(n - 1)
		return sum
	}
	_ = n > 0 && f() > 0
	return sum
}




// 不能有 for
// func bitAdd(a,b int) int {

//     for b != 0 {
//         sum := a ^ b
//         carry := (a & b) << 1
//         a = sum
//         b = carry
//     }

//     return a
// }
