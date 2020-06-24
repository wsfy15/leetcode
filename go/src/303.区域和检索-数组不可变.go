/*
 * @lc app=leetcode.cn id=303 lang=golang
 *
 * [303] 区域和检索 - 数组不可变
 */

// @lc code=start
type NumArray struct {
	nums []int
	sum []int
}


func Constructor(nums []int) NumArray {
	na := NumArray{ nums: nums }
	n := len(nums)
	if n == 0 {
		return na
	}

	sum := make([]int, n)
	sum[0] = nums[0]
	for i := 1; i < n; i++ {
		sum[i] = sum[i - 1] + nums[i]
	}

	na.sum = sum
	return na
}


func (this *NumArray) SumRange(i int, j int) int {
	if len(this.nums) == 0 {
		return 0
	}
	
	if i == 0 {
		return this.sum[j]
	}
	return this.sum[j] - this.sum[i - 1]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
// @lc code=end

