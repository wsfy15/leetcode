/*
 * @lc app=leetcode.cn id=213 lang=golang
 *
 * [213] 打家劫舍 II
 */

// @lc code=start
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n < 4 {
		return max(nums...)
	} 
	return max(helper(nums, 0, n - 2), helper(nums, 1, n - 1))
}

func helper(nums []int, start, end int) int {
	a, b := nums[start], 0 // a：偷 b：不偷
	for i := start + 1; i <= end; i++ {
		a, b = b + nums[i], max(a, b)
	}
	return max(a, b)
}

func max(a ...int) int {
	res := a[0]
	for _, v := range a {
		if res < v {
			res = v
		}
	}
	return res
}
// @lc code=end

