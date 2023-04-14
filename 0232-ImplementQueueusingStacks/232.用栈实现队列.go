/*
 * @lc app=leetcode.cn id=232 lang=golang
 *
 * [232] 用栈实现队列
 */

// @lc code=start
type MyQueue struct {
	inStack  []int
	outStack []int
}

func Constructor() MyQueue {
	return MyQueue{
		inStack:  make([]int, 0, 128),
		outStack: make([]int, 0, 128),
	}
}

func (this *MyQueue) Push(x int) {
	this.inStack = append(this.inStack, x)
}

func (this *MyQueue) transInStackToOutStack() {
	if len(this.outStack) != 0 {
		// outStack不为空时禁止转移
		return
	}

	for len(this.inStack) > 0 {
		this.outStack = append(this.outStack, this.inStack[len(this.inStack)-1])
		this.inStack = this.inStack[0 : len(this.inStack)-1]
	}
}

func (this *MyQueue) Pop() int {
	if len(this.outStack) == 0 {
		// inStack => outStack
		this.transInStackToOutStack()
	}
	val := this.outStack[len(this.outStack)-1]
	this.outStack = this.outStack[0 : len(this.outStack)-1]
	return val
}

func (this *MyQueue) Peek() int {
	if len(this.outStack) == 0 {
		// inStack => outStack
		this.transInStackToOutStack()
	}
	return this.outStack[len(this.outStack)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.inStack) == 0 && len(this.outStack) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// @lc code=end

