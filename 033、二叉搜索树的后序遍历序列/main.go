func verifyPostorder(postorder []int) bool {

    if len(postorder) <= 2 {
        return true
    }

    root := postorder[len(postorder) - 1]

    for i, v := range postorder[:len(postorder) - 1] {
        if v < root {
            continue
        }else {
            left := postorder[:i]
            right := postorder[i:len(postorder) - 1]

            fmt.Println(left,i,right)
            for _,v := range left {
                if v > root {
                    return false
                }
            }

            for _,v := range right {
                if v < root {
                    return false
                }
            }

            return verifyPostorder(left) && verifyPostorder(right)
        }
    }



    return verifyPostorder(postorder[:len(postorder) - 1])

}