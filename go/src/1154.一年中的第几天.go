/*
 * @lc app=leetcode.cn id=1154 lang=golang
 *
 * [1154] 一年中的第几天
 */

// @lc code=start
func dayOfYear(date string) int {
	months := []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	items := strings.Split(date, "-")
	year, _ := strconv.Atoi(items[0])
	month, _ := strconv.Atoi(items[1])
	day, _ := strconv.Atoi(items[2])

	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		months[2]++
	}

	res := day
	for i := 1; i < month; i++ {
		res += months[i]
	}
	return res
}

// @lc code=end

