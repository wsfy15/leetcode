/*
 * @lc app=leetcode.cn id=275 lang=golang
 *
 * [275] H 指数 II
 */

// @lc code=start
func hIndex(citations []int) int {
	n := len(citations)
	l, r := 0, n
	for l < r {
		mid := l + (r-l)>>1

		if citations[mid] < n-mid {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return n - l
}

// @lc code=end

