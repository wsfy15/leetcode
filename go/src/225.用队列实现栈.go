/*
 * @lc app=leetcode.cn id=225 lang=golang
 *
 * [225] 用队列实现栈
 */

// @lc code=start
type MyStack struct {
	masterQueue []int
	slaveQueue []int
}


/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	this.masterQueue = append(this.masterQueue, x)
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	for len(this.masterQueue) > 1 {
		this.slaveQueue = append(this.slaveQueue, this.masterQueue[0])
		this.masterQueue = this.masterQueue[1:]
	}
	val := this.masterQueue[0]
	this.masterQueue, this.slaveQueue = this.slaveQueue, make([]int, 0)
	return val
}


/** Get the top element. */
func (this *MyStack) Top() int {
	for len(this.masterQueue) > 1 {
		this.slaveQueue = append(this.slaveQueue, this.masterQueue[0])
		this.masterQueue = this.masterQueue[1:]
	}
	return this.masterQueue[0]
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.masterQueue) == 0 && len(this.slaveQueue) == 0
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

