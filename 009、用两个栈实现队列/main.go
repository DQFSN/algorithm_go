type CQueue struct {
    input Stack
    output Stack
}


func Constructor() CQueue {
    q :=CQueue{}
    var s Stack
    q.input = s
    var s2 Stack
    q.output = s2

    return q
}


func (this *CQueue) AppendTail(value int)  {
    this.input.Push(value)
}


func (this *CQueue) DeleteHead() int {
    if len(this.output) == 0 {
        for len(this.input) != 0 {
            ret, err := this.input.Pop()
            if err == nil {
                this.output.Push(ret)
            }else {
                break
            }
        }
    }

    ret, err := this.output.Pop()
    if err != nil {
        return -1
    }else {
        return ret.(int)
    }

}


/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */


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
