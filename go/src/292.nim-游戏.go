/*
 * @lc app=leetcode.cn id=292 lang=golang
 *
 * [292] Nim 游戏
 */

// @lc code=start
// 如果能留给对手4个，就能赢
// 5、6、7  9、10、11 ……
// 每次拿完总能留给对手4n个
func canWinNim(n int) bool {
	// return n % 4 != 0
	return n & 3 > 0
}
// @lc code=end

