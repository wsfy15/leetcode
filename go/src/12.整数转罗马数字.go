/*
 * @lc app=leetcode.cn id=12 lang=golang
 *
 * [12] 整数转罗马数字
 */

// @lc code=start
func intToRoman(num int) string {
	value := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	str := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	sb := strings.Builder{}
	for i := 0; i < len(value) && num > 0; i++ {
		for num >= value[i] {
			num -= value[i]
			sb.WriteString(str[i])
		}
	}
	// for num > 0 {
	// 	for i := 0; i < len(value); i++ {
	// 		if num >= value[i] {
	// 			num -= value[i]
	// 			sb.WriteString(str[i])
	// 			break
	// 		}
	// 	}
	// }

	return sb.String()
}

// @lc code=end

