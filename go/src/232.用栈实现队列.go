/*
 * @lc app=leetcode.cn id=232 lang=golang
 *
 * [232] 用栈实现队列
 */

// @lc code=start
type MyQueue struct {
		input []int
		output []int	// 每当有peek或者top操作时触发从input到output的搬迁（output为空的条件下）
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
    return MyQueue{}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
    this.input = append(this.input, x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.output) == 0 {
		this.move()
	}
	val := this.output[len(this.output) - 1]
	this.output = this.output[:len(this.output) - 1]
	return val
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.output) == 0 {
		this.move()
	}
	return this.output[len(this.output) - 1]
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
    return len(this.input) == 0 && len(this.output) == 0
}

func (this *MyQueue) move() {
	for i := len(this.input) - 1; i >= 0; i-- {
		this.output = append(this.output, this.input[i])
	}
	this.input = append([]int{})
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

