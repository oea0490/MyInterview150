package Stack

import "math"

type MinStack struct {
	xStack   []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		xStack:   []int{},
		minStack: []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.xStack = append(this.xStack, val)
	this.minStack = append(this.minStack, min(val, this.GetMin()))
}

func (this *MinStack) Pop() {
	this.xStack = this.xStack[:len(this.xStack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.xStack[len(this.xStack)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.minStack) == 0 {
		return math.MaxInt64
	}
	return this.minStack[len(this.minStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
