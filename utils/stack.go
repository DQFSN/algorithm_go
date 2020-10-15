package utils

import "errors"

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
