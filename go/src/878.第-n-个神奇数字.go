/*
 * @lc app=leetcode.cn id=878 lang=golang
 *
 * [878] 第 N 个神奇数字
 */

// @lc code=start
func nthMagicalNumber(n int, a int, b int) int {
	var mod int = 1e9 + 7
	var lcm int64 = int64(a * b / gcd(a, b))

	var l, h int64 = 0, 1e15
	for l < h {
		mid := l + (h-l)>>1
		if mid/int64(a)+mid/int64(b)-mid/lcm < int64(n) {
			l = mid + 1
		} else {
			h = mid
		}
	}

	return int(l) % mod
}

func gcd(a, b int) int {
	if a < b {
		return gcd(b, a)
	}

	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// @lc code=end

