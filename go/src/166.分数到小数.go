/*
 * @lc app=leetcode.cn id=166 lang=golang
 *
 * [166] 分数到小数
 */

// @lc code=start
func fractionToDecimal(numerator int, denominator int) string {
	var res string
	if numerator * denominator < 0 {
		res = "-"
	}

	numerator, denominator = abs(numerator), abs(denominator)
	
	res += strconv.Itoa(numerator / denominator)
	numerator %= denominator
	if numerator == 0 {
		return res
	}

	res += "."
	m := map[int]int{}
	var str string
	for numerator > 0 {
		if v, ok := m[numerator]; ok {
			res += str[:v]
			res += "(" + str[v:] + ")"
			return res
		}

		m[numerator] = len(str)
		numerator *= 10

		str += strconv.Itoa(numerator / denominator)
		numerator = numerator % denominator
	} 

	if numerator == 0 {
		res += str
	}
	return res
}

func abs(x int) int {
	y := x >> 31
	return (x ^ y) - y
}
// @lc code=end

