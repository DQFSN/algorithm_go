type MedianFinder struct {
    nums []int
    size int
}


/** initialize your data structure here. */
func Constructor() MedianFinder {
    return MedianFinder{
        nums:make([]int,50000,50000),
    }
}


func (this *MedianFinder) AddNum(num int)  {
     if this.size == 0 {
         this.nums[this.size] = num
     }else {

        index:=this.size-1
        for ; index>=0 && this.nums[index]>num; index-- {
            this.nums[index+1] = this.nums[index]
        }
        this.nums[index+1] = num
     }
    this.size++

}


func (this *MedianFinder) FindMedian() float64 {
    if this.size % 2 == 0 {
        return float64(this.nums[this.size/2] + this.nums[this.size/2 - 1]) / 2.0
    }else {
        return float64(this.nums[this.size/2])
    }
}


/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */