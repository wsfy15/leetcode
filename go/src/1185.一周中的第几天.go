/*
 * @lc app=leetcode.cn id=1185 lang=golang
 *
 * [1185] 一周中的第几天
 */

// @lc code=start
func dayOfTheWeek(day int, month int, year int) string {
	if month < 3 {
		month += 12
		year--
	}

	// 基姆拉尔森计算公式
	index := (day + 2*month + 3*(month+1)/5 + year + year/4 - year/100 + year/400) % 7
	weekdays := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	return weekdays[index]
}

// @lc code=end

