/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 */

// @lc code=start
type MinStack struct {
		vals []int
		count int

		mins []int	// 存放每个阶段的最小值
		minCount int 
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{
			count: 0,
			minCount: 0,
		}
}


func (this *MinStack) Push(x int)  {
		this.vals = append(this.vals, x)
		if len(this.mins) == 0 || x <= this.mins[this.minCount - 1] {
			this.mins = append(this.mins, x)
			this.minCount++
		}
		this.count++
}


func (this *MinStack) Pop()  {
	 min := this.vals[this.count - 1]
	 this.vals = this.vals[:this.count - 1]
	 this.count--

	 // 出栈的元素是当前最小值
	 if min == this.mins[this.minCount - 1] {
		this.mins = this.mins[:this.minCount - 1]
		this.minCount--
	 }
}


func (this *MinStack) Top() int {
    return this.vals[this.count - 1]
}


func (this *MinStack) GetMin() int {
    return this.mins[this.minCount - 1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end

