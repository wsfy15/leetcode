/*
 * @lc app=leetcode.cn id=492 lang=golang
 *
 * [492] 构造矩形
 */

// @lc code=start
func constructRectangle(area int) []int {
	res := int(math.Sqrt(float64(area)))
	for area % res != 0 {
		res--
	}

	if res > area / res {
		return []int{res, area / res}
	}
	return []int{area / res, res}
}
// @lc code=end

