/*
 * @lc app=leetcode.cn id=575 lang=golang
 *
 * [575] 分糖果
 */

// @lc code=start
func distributeCandies(candies []int) int {
	n := len(candies)
	m := map[int]int{}
	for _, candy := range candies {
		m[candy]++
	}

	if len(m) > n >> 1 {
		return n >> 1
	}
	return len(m)
}
// @lc code=end

