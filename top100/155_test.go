package top100

type pair struct {
	val    int
	preMin int // 需要记录一下元素的前一最小值，方便pop的时候，知道哪一个最小
}

type MinStack struct {
	mp []pair
}

func Constructor() MinStack {
	return MinStack{mp: []pair{{}}}
}

func (this *MinStack) Push(val int) {
	this.mp = append(this.mp, pair{val, min(val, this.GetMin())})
}

func (this *MinStack) Pop() {
	this.mp = this.mp[:len(this.mp)-1]
}

func (this *MinStack) Top() int {
	return this.mp[len(this.mp)-1].val
}

// 最后一个元素push进来的时候preMin即为最小值
func (this *MinStack) GetMin() int {
	if len(this.mp) <= 0 {
		return 0
	}
	return this.mp[len(this.mp)-1].preMin
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
