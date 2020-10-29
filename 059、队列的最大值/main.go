type MaxQueue struct {
    queue []int
    max int
    maxQueue []int
}


func Constructor() MaxQueue {
    return MaxQueue{}
}


func (this *MaxQueue) Max_value() int {
    if len(this.maxQueue) == 0 {
        return -1
    }

    return this.maxQueue[0]
}


func (this *MaxQueue) Push_back(value int)  {
    this.queue = append(this.queue,value)

    if value > this.max {
        this.max = value

        for i := range this.maxQueue {
            this.maxQueue[i] = value
        }
    }else {
        for i := range this.maxQueue {
            if this.maxQueue[i] < value {
                this.maxQueue[i] = value
            }
        }
    }

    this.maxQueue = append(this.maxQueue,value)

}


func (this *MaxQueue) Pop_front() int {

    if len(this.queue) == 0{
        return -1 
    }

    res := this.queue[0]
    this.queue = this.queue[1:]
    this.maxQueue = this.maxQueue[1:]
    this.max = this.Max_value()

    return res
}


/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */