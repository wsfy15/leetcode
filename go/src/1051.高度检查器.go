/*
 * @lc app=leetcode.cn id=1051 lang=golang
 *
 * [1051] 高度检查器
 */

// @lc code=start
func heightChecker(heights []int) int {
	tmp := make([]int, len(heights))
	copy(tmp, heights)
	sort.Ints(tmp)
	res := 0
	for i := 0; i < len(heights); i++ {
		if tmp[i] != heights[i] {
			res++
		}
	}
	return res
}
// @lc code=end

