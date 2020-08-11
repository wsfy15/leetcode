/*
 * @lc app=leetcode.cn id=1323 lang=golang
 *
 * [1323] 6 和 9 组成的最大数字
 */

// @lc code=start
func maximum69Number(num int) int {
	str := strconv.Itoa(num)
	bytes := []byte(str)
	for i := 0; i < len(bytes); i++ {
		if bytes[i] == '6' {
			bytes[i] = '9'
			break
		}
	}
	res, _ := strconv.Atoi(string(bytes))
	return res
}

// @lc code=end

