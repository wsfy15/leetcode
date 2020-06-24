/*
 * @lc app=leetcode.cn id=541 lang=golang
 *
 * [541] 反转字符串 II
 */

// @lc code=start
func reverseStr(s string, k int) string {
	bytes := []byte(s)
	n, i := len(s), 0
	for i < n {
		left, right := i, i + k - 1
		if right >= n {
			right = n - 1
		}
		for left < right {
			bytes[left], bytes[right] = bytes[right], bytes[left]
			left++
			right--
		}

		i += 2 * k
	}

	return string(bytes)
}
// @lc code=end

