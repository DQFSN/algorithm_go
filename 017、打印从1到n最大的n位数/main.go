package main

import "fmt"

func main()  {
	printNumbers(3)
}
// // 脱裤子放屁
// func printNumbers(n int) []int {
//     // numByte := make([]byte, n, n)

//     maxLen := int(myPow(10, n))
//     result := make([]int, maxLen, maxLen)

    
//     for i := 1; i <= n; i++ {
//         preLen := int(myPow(10,i))
//         pre := result[:preLen]

//         if i == 1 {
//             pre[0] = 0
//             pre[1] = 1
//             pre[2] = 2
//             pre[3] = 3
//             pre[4] = 4
//             pre[5] = 5
//             pre[6] = 6
//             pre[7] = 7
//             pre[8] = 8
//             pre[9] = 9
//         }
        
//         for base,index:=1,preLen; base <= 9 && index < maxLen; base++ {
//             index = preLen * base
//             for _, v:= range pre {

//                 result[index] = preLen*base + v
//                 index++ 
//             }
//         }
//     }

//     return result[1:]
// }

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



// 大数版,字符串加法

func printNumbers(n int) {
	max := int(myPow(10,n))
	result := make([]byte,max,max)

	for max > 1 {
		addOne(result)
		temp := result[numIndex(result):]
		
		for _,v := range temp {
			fmt.Printf("%d", v)
		}
		fmt.Printf("\t")
		max--
	}
	fmt.Printf("\n")


}

func addOne(num []byte) {
	l := len(num)

	var carry byte = 0
	step := 1
	for i:=l - 1;i >= 0;i-- {
		a := num[i]
		b := a + carry + byte(step)
		carry = b / 10
		num[i] = b % 10
		step = 0
	}

}

// 第一个非零‘0’下标
func numIndex(num []byte) int {

	for i,v:=range num {
		if v != 0 {
			return i 
		}
	}

    return len(num)
}


