/*
 * @lc app=leetcode.cn id=344 lang=golang
 *
 * [344] 反转字符串
 */

// @lc code=start
func reverseString(s []byte)  {
	size := len(s)
	if size <= 1 {
		return
	}

	left := 0
	right := size - 1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
// @lc code=end

