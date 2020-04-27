/*
 * @lc app=leetcode.cn id=367 lang=golang
 *
 * [367] 有效的完全平方数
 */

// @lc code=start
func isPerfectSquare(num int) bool {
	low, high := 1, num
	for low < high {
		mid := low + (high - low) >> 1

		tmp := mid * mid 
		if tmp == num {
			return true
		}

		if tmp < num {
			if (mid+1) * (mid+1) > num {
				return false
			}
			low = mid + 1
		} else {
			high = mid
		}
	}

	return low * low == num
}
// @lc code=end

