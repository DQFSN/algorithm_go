package main

import "fmt"

var (
    cMax, cMin = 0, 0
    rMax, rMin = 0, 0
)

type point struct {
    r, c int
}




func spiralOrder(matrix [][]int) []int {
    rNum := len(matrix)
    if rNum == 0 {
        return nil
    }
    cNum := len(matrix[0])

    result := make([]int, rNum * cNum, rNum * cNum)

    cMax, cMin = cNum - 1, 0
    rMax, rMin = rNum - 1, 0

    posi := point{r : 0, c : 0}
    opts := []point{
        {r : 0, c : 1},
        {r : 1, c : 0},
        {r : 0, c : -1},
        {r : -1, c : 0},
    }

	index := 0 
	corner := false
	optLoop := 0
    for  {

        if index == cNum*rNum {
            break
        }
		opt := opts[optLoop % 4]
		optLoop++

        if rMin > rMax && cMin > cMax {
            break
        }

        for !changeOpt(posi) || corner {

            if corner == false {
				result[index] = matrix[posi.r][posi.c]
            	index++
			}
            posi.r += opt.r
			posi.c += opt.c
			corner = false
        }

		corner = true

        switch opt {
            case point{r : 0, c : 1}:
                posi.c--
                rMin++
            case point{r : 1, c : 0}:
                posi.r--
                cMax--
            case point{r : 0, c : -1}:
                posi.c++
                rMax--
            case point{r : -1, c : 0}:
                posi.r++
                cMin++
        }
    } 

    return result
}

func changeOpt(posi point) bool {
    if posi.r < rMin || posi.r > rMax || posi.c < cMin || posi.c > cMax {
        return true
    }
    return false
}

func main()  {
	
	// a := [][]int{{1,2,3,4},{12,13,14,5},{11,16,15,6},{10,9,8,7}}
	a := [][]int{{1,2,3},{4,5,6},{7,8,9}}
	// a := [][]int{{1,2},{4,3}}
	// a := [][]int{{1,2}}
	// a := [][]int{{1}}
	fmt.Println(spiralOrder(a))
}