/*
 * @lc app=leetcode.cn id=1362 lang=golang
 *
 * [1362] 最接近的因数
 */

// @lc code=start
func closestDivisors(num int) []int {
	var res []int
	start := int(math.Sqrt(float64(num + 2)))
	for start > 0 {
		if (num+1)%start == 0 {
			return []int{start, (num + 1) / start}
		}
		if (num+2)%start == 0 {
			return []int{start, (num + 2) / start}
		}
		start--
	}
	return res
}

// @lc code=end

