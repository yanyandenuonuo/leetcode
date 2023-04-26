/*
 * @lc app=leetcode.cn id=225 lang=golang
 *
 * [225] 用队列实现栈
 */

// @lc code=start
type MyStack struct {
	inQueue  []int
	outQueue []int
}

func Constructor() MyStack {
	return MyStack{
		inQueue:  []int{},
		outQueue: []int{},
	}
}

func (this *MyStack) Push(x int) {
	for idx := range this.outQueue {
		this.inQueue = append(this.inQueue, this.outQueue[idx])
	}
	this.outQueue = []int{x}
}

func (this *MyStack) Pop() int {
	if this.Empty() {
		return 0
	}
	if len(this.outQueue) == 1 {
		topVal := this.outQueue[0]
		this.outQueue = []int{}
		return topVal
	}
	topVal := 0
	for idx := 0; idx < len(this.inQueue); idx += 1 {
		topVal = this.inQueue[idx]
		this.outQueue = append(this.outQueue, this.inQueue[idx])
	}
	this.inQueue = []int{}
	if len(this.outQueue) >= 2 {
		for idx := 0; idx < len(this.outQueue)-2; idx += 1 {
			this.inQueue = append(this.inQueue, this.outQueue[idx])
		}
		this.outQueue = []int{this.outQueue[len(this.outQueue)-2]}
	} else {
		this.outQueue = []int{}
	}

	return topVal
}

func (this *MyStack) Top() int {
	if this.Empty() {
		return 0
	}
	if len(this.outQueue) == 1 {
		return this.outQueue[0]
	}
	topVal := 0
	for idx := 0; idx < len(this.inQueue); idx += 1 {
		topVal = this.inQueue[idx]
		this.outQueue = append(this.outQueue, this.inQueue[idx])
	}
	this.inQueue = []int{}
	for idx := 0; idx < len(this.outQueue)-1; idx += 1 {
		this.inQueue = append(this.inQueue, this.outQueue[idx])
	}
	this.outQueue = []int{this.outQueue[len(this.outQueue)-1]}

	return topVal
}

func (this *MyStack) Empty() bool {
	return len(this.inQueue) == 0 && len(this.outQueue) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
// @lc code=end

