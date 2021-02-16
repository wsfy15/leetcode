/*
 * @lc app=leetcode.cn id=1208 lang=golang
 *
 * [1208] 尽可能使字符串相等
 */

// @lc code=start
func equalSubstring(s string, t string, maxCost int) int {
	n := len(s)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = abs(int(s[i]) - int(t[i]))
	}

	l, r, cur, res := 0, 0, 0, 0
	for r < n {
		cur += nums[r]
		for cur > maxCost {
			cur -= nums[l]
			l++
		}

		if r - l + 1 > res {
			res = r - l + 1
		}

		r++
	}
	return res
}

func abs(x int) int {
	y := x >> 31
	return (x ^ y) - y
}
// @lc code=end

