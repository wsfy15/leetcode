/*
 * @lc app=leetcode.cn id=796 lang=golang
 *
 * [796] 旋转字符串
 */

// @lc code=start
func rotateString(A string, B string) bool {
	return len(A) == len(B) && strings.Contains(A+A, B)
}
// @lc code=end

