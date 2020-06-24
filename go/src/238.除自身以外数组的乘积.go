/*
 * @lc app=leetcode.cn id=238 lang=golang
 *
 * [238] 除自身以外数组的乘积
 */

// @lc code=start
func productExceptSelf(nums []int) []int {
	// 左右累乘
	n := len(nums)
	res := make([]int, n)
	left, right := 1, 1
	for i := 0; i < n; i++ {
		res[i] = left
		left *= nums[i]
	}

	for i := n - 1; i >= 0; i-- {
		res[i] *= right
		right *= nums[i]
	}

	return res
}
// @lc code=end

