/*
 * @lc app=leetcode.cn id=933 lang=golang
 *
 * [933] 最近的请求次数
 */

// @lc code=start
type RecentCounter struct {
	queue []int
}


func Constructor() RecentCounter {
	return RecentCounter{}
}


func (this *RecentCounter) Ping(t int) int {
	min := t - 3000
	for len(this.queue) > 0 && this.queue[0] < min {
		this.queue = this.queue[1:]
	}
	this.queue = append(this.queue, t)
	return len(this.queue)
}


/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
// @lc code=end

