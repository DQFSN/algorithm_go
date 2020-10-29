type Point struct {
    x,y int
}


func twoSum(n int) []float64 {
    pro := make([]float64, 5*n+1, 5*n+1)

    dict := make(map[Point]int)
    dict[Point{1,1}] = 1
    dict[Point{1,2}] = 1
    dict[Point{1,3}] = 1
    dict[Point{1,4}] = 1
    dict[Point{1,5}] = 1
    dict[Point{1,6}] = 1
    all := 0.0
    sumEuqalsS(n,6*n,&dict)
    for i := range pro {
        pro[i] = float64(dict[Point{n,n+i}])
        // fmt.Println(pro)
        all += pro[i]
    }
    for i := range pro {
        pro[i] /= all
    }


    return pro
}

func sumEuqalsS(n, s int, dict *map[Point]int) {

    // res := 0
    // if n == 1 {
    //     if s == 1 || s == 2 || s == 3 || s == 4 || s == 5 || s == 6 {
    //         return 1
    //     }
    //     return 0
    // }

    // res = sumEuqalsS(n-1,s-1)+sumEuqalsS(n-1,s-2)+sumEuqalsS(n-1,s-3)+sumEuqalsS(n-1,s-4)+sumEuqalsS(n-1,s-5)+sumEuqalsS(n-1,s-6)

    for c :=2;c<=n;c++ {
        for sum:=1*c;sum<=6*c;sum++ {
            // sum := (c - 1)*6 + i
            (*dict)[Point{c,sum}] = (*dict)[Point{c-1,sum-1}]+(*dict)[Point{c-1,sum-2}]+(*dict)[Point{c-1,sum-3}]+(*dict)[Point{c-1,sum-4}]+(*dict)[Point{c-1,sum-5}]+(*dict)[Point{c-1,sum-6}]
        }
    }

}