/*
 * @lc app=leetcode.cn id=860 lang=golang
 *
 * [860] 柠檬水找零
 */

// @lc code=start
func lemonadeChange(bills []int) bool {
	five, ten := 0, 0

	for _, num := range bills {
		switch num {
		case 5:
			five++
		case 10:
			if five <= 0 {
				return false
			}
			five--
			ten++
		case 20:
			left := 15
			if ten > 0 {
				ten--
				left = 5
			}
			five -= left / 5
			if five < 0 {
				return false
			}
		}
	}

	return true
}

// @lc code=end

