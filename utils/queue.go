package utils

import "errors"

type Queue []interface{}

func (q *Queue) Push(v interface{})  {
	*q = append(*q, v)
}

func (q *Queue) Pop() (interface{}, error) {

	if len(*q) == 0 {
		return nil, errors.New("no data")
	}

	ret := (*q)[0]
	*q = (*q)[1:]

	return ret, nil
}


func (q *Queue) Front() (interface{}, error) {

	if len(*q) == 0 {
		return nil, errors.New("no data")
	}

	ret := (*q)[0]

	return ret, nil
}