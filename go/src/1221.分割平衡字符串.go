/*
 * @lc app=leetcode.cn id=1221 lang=golang
 *
 * [1221] 分割平衡字符串
 */

// @lc code=start
func balancedStringSplit(s string) int {
	res := 0
	l, r := 0, 0
	for _, ch := range []byte(s) {
		if ch == 'L' {
			l++
		} else {
			r++
		}
		if l == r {
			l = 0
			r = 0
			res++
		}
	}
	return res
}

// @lc code=end

