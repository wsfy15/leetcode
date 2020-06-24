/*
 * @lc app=leetcode.cn id=507 lang=golang
 *
 * [507] 完美数
 */

// @lc code=start
func checkPerfectNumber(num int) bool {
	if num < 2 {
		return false
	}

	total, cur := 1, 2
	upper := int(math.Sqrt(float64(num)))
	for cur <= upper {
		if num % cur == 0 {
			total += cur + num / cur
			if total > num {
				return false
			}
		}
		cur++
	} 
	return total == num
}
// @lc code=end

