/*
 * @lc app=leetcode.cn id=1360 lang=golang
 *
 * [1360] 日期之间隔几天
 */

// @lc code=start
func daysBetweenDates(date1 string, date2 string) int {
	return abs(getDays(date1) - getDays(date2))
}

// 返回从1971年1月1日到date的天数
func getDays(date string) int {
	words := strings.Split(date, "-")
	year, _ := strconv.Atoi(words[0])
	month, _ := strconv.Atoi(words[1])
	day, _ := strconv.Atoi(words[2])

	res := 0
	monthDays := []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	for i := 1971; i < year; i++ {
		res += 365 + isLeap(i)
	}
	for i := 1; i < month; i++ {
		if i == 2 {
			res += isLeap(year)
		}
		res += monthDays[i]
	}
	return res + day
}

func isLeap(n int) int {
	if n%4 == 0 && (n%100 != 0 || n%400 == 0) {
		return 1
	}
	return 0
}

func abs(x int) int {
	y := x >> 31
	return (x ^ y) - y
}

// @lc code=end

