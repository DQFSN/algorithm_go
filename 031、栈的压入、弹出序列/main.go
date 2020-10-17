type Stack []interface{}

func (s *Stack) Push(v interface{}) {
	*s = append(*s, v)
}

func (s *Stack) Pop() (ret interface{}, err error) {

	if len(*s) == 0{
		return nil, errors.New("no data")
	}

	ret = (*s)[len(*s) - 1]
	*s = (*s)[:len(*s) - 1]
	return ret, nil
}


func (s *Stack) Peak() (ret interface{}, err error) {
	if len(*s) == 0 {
		return nil, errors.New("no data")
	}

	ret = (*s)[len(*s) - 1]

	return ret, nil
}

// func validateStackSequences2(pushed []int, popped []int) bool {
//     var stack Stack

//     for _, v := range pushed {
//         stack.Push(v)
//         top,_ := stack.Peak()
//         for len(popped) > 0 && top == popped[0] {
//             stack.Pop()
//             top,_ = stack.Peak()
//             popped = popped[1:]
//         }
//     }

//     return len(stack) == len(popped) && len(popped) == 0
// }



func validateStackSequences(pushed []int, popped []int) bool {
    // cap 影响时间很明显
    stack := make([]int, 0, 1000)

    for _, v := range pushed {
        stack = append(stack,v)
        top := stack[len(stack) - 1]
        for len(popped) > 0 && top == popped[0] {
            stack = stack[:len(stack) - 1]
            popped = popped[1:]
            if len(stack) == 0 {
                break
            }
            top = stack[len(stack) - 1]
            
        }
    }

    return len(stack) == len(popped) && len(popped) == 0
}

