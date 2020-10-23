package main

import "fmt"

func main()  {
	//var s utils.Stack
	//s.Push(1)
	//fmt.Println(s.Peak())
	//s.Pop()
	//fmt.Println(s.Peak())

	s := "sdasdsa"
	b := []byte(s)
	bb := b

	fmt.Printf("%p\t%p\t%p", &s, &b, &bb)
}