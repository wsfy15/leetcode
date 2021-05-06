/*
 * @lc app=leetcode.cn id=1247 lang=golang
 *
 * [1247] 交换字符使得字符串相同
 */

// @lc code=start
func minimumSwap(s1 string, s2 string) int {
	n := len(s1)
	x, y := 0, 0
	for i := 0; i < n; i++ {
		if s1[i] == s2[i] {
			continue
		}

		if s1[i] == 'x' {
			x++
		} else {
			y++
		}
	}

	if x != y && (x+y)%2 == 1 {
		return -1
	}

	res := x/2 + y/2
	res += 2 * (x % 2)
	return res
}

// @lc code=end

