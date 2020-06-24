/*
 * @lc app=leetcode.cn id=657 lang=golang
 *
 * [657] 机器人能否返回原点
 */

// @lc code=start
func judgeCircle(moves string) bool {
	upCount, rightCount := 0, 0
	for i := 0; i < len(moves); i++ {
		switch moves[i] {
		case 'R':
			rightCount++
		case 'L':
			rightCount--
		case 'U':
			upCount++
		case 'D':
			upCount--
		}
	}
	return upCount == 0 && rightCount == 0
}
// @lc code=end

