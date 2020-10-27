package _65_不用加减乘除做加法



func add(a int, b int) int {


	for b != 0 {
		sum := a ^ b
		carry := (a & b) << 1
		a = sum
		b = carry
	}

	return a
}

