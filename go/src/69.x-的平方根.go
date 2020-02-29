/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 */

// @lc code=start
func mySqrt(x int) int {
    if x < 2 {
			return x
		}

		low := 1
		high := x - 1
		for low <= high {
			mid := low + ((high - low) >> 1)
			if mid * mid <= x {
				if (mid + 1) * (mid + 1) > x {
					return mid
				} else {
					low = mid + 1
				}
			} else {
				high = mid - 1
			}
		}
		return low
}
// @lc code=end

