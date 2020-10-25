func cuttingRope(n int) int {

    switch n {
        case 1,2:
            return 1
        case 3:
            return 2
    }

    mod := n % 3
    numThree := n / 3

    result := 0

    switch mod {
        case 0:
            result = pow(3, numThree)
        case 1:
            result = pow(3, numThree - 1) * 4
        case 2:
            result = pow(3, numThree) * 2
    }

    return result
}

func pow(base, exp int) int {

    result := 1

    for exp != 0 {
        if exp & 1 == 1{
            result = result * base
        }
        exp = exp >> 1
        base = base * base
    }

    return result

}


// 下面是大数版本



func cuttingRope(n int) int {

    switch n {
        case 1,2:
            return 1
        case 3:
            return 2
    }

    mod := n % 3
    numThree := n / 3

    result := 0

    switch mod {
        case 0:
            result = pow(3, numThree)
        case 1:
            result = pow(3, numThree - 1) * 4
        case 2:
            result = pow(3, numThree) * 2
    }

    return result % 1000000007
}

func pow(base, exp int) int {

    result := 1

    for exp != 0 {
        if exp & 1 == 1{
            result = (result * base) % 1000000007
        }
        exp = exp >> 1
        base = (base * base) % 1000000007
    }

    return result

}