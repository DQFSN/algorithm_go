func maxProfit(prices []int) int {

    if len(prices) <= 1 {
        return 0
    }

    min := prices[0]
    maxPro := 0

    for _, v := range prices[1:] {
        if v - min > maxPro {
            maxPro = v - min
        }
        if v < min {
            min = v
        }
    }

    return maxPro
}