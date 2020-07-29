/*
 * @lc app=leetcode.cn id=453 lang=golang
 *
 * [453] 最小移动次数使数组元素相等
 */

// @lc code=start
func minMoves(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}

	// 每次让n-1个数加一，相当于让1个数减1
	num, total := min(nums...), 0
	for _, v := range nums {
		total += v - num
	}
	return total
}

func min(a ...int) int {
	res := a[0]
	for _, v := range a {
		if res > v {
			res = v
		}
	}
	return res
}
// @lc code=end

