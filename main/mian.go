package main

import "algorithm_go/utils"
import "fmt"

func main()  {
	var s utils.Stack
	s.Push(1)
	fmt.Println(s.Peak())
	s.Pop()
	fmt.Println(s.Peak())
}