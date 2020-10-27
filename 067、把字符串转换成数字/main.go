func strToInt(str string) int {

	str = strings.Trim(str, " ")

	if len(str) == 0 {
		return 0
	}

	neg := false
	var num int32 = 0
	if  '0' <= str[0] && str[0] <= '9'  {
		num = parseInt(str)
	}else if str[0] == '-' {
		neg = true
		num = parseInt(str[1:])
	}else if str[0] == '+' {
		num = parseInt(str[1:])
	}else {
		return 0
	}

	if neg {
		if num == -1 {
			return -2147483648
		}
		return int(-num)
	}else {
		if num == -1 {
			return 2147483647
		}
		return int(num)
	}

}

func parseInt(str string) int32 {
var num int32 = 0

// overflow ：= false
for _, v := range str {

	if v < '0' || v > '9' {
		break
	}

	pre := num
	num = num*10 + (v - '0') 
	// 溢出 单纯 num < pre 不正确
	if (num - (v - '0'))/10 != pre || pre > num{
		// overflow = true
		return -1
	}
}

return num

}