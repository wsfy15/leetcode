/*
 * @lc app=leetcode.cn id=453 lang=golang
 *
 * [453] 最小移动次数使数组元素相等
 */

// @lc code=start
func minMoves(nums []int) int {
	n := len(nums)
	max, total := math.MinInt32, 0 
	for _, v := range nums {
		total += v
		if v > max {
			max = v
		}
	}

	for {
		if (n * max - total) % (n - 1) == 0 {
			return (n * max - total) / (n - 1)
		}
		max++
	}
}
// @lc code=end

