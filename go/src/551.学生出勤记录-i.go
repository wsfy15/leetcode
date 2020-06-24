/*
 * @lc app=leetcode.cn id=551 lang=golang
 *
 * [551] 学生出勤记录 I
 */

// @lc code=start
func checkRecord(s string) bool {
	absent, late := 0, 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'A':
			late = 0
			absent++
			if absent > 1 {
				return false
			}
		case 'L':
			late++
			if late > 2 {
				return false
			}
		case 'P':
			late = 0
		}
	}
	return true
}
// @lc code=end

