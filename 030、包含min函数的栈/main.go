type MinStack struct {
    data []int
    m []int
    min int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{min:0}
}


func (this *MinStack) Push(x int)  {

    if x < this.min || len(this.m) == 0  {
        this.min = x
    }

    this.data = append(this.data, x)
    this.m = append(this.m, this.min)
}


func (this *MinStack) Pop()  {
    if len(this.data) == 0 {
        return
    }
    this.data = this.data[:len(this.data) - 1]
    this.m = this.m[:len(this.m) - 1]
    this.min = this.Min()
}


func (this *MinStack) Top() int {
    return this.data[len(this.data) - 1]
}


func (this *MinStack) Min() int {
    if len(this.m) == 0 {
        return 0
    }
    return this.m[len(this.m) - 1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */