/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 */

// @lc code=start
func romanToInt(s string) int {
	m := map[byte]int{
		'I': 1,
		// 'IV': 4,
		'V': 5,
		// 'IX': 9,
		'X': 10,
		// 'XL': 40,
		'L': 50,
		// 'XC': 90,
		'C': 100,
		// 'CD': 400,
		'D': 500,
		// 'CM': 900,
		'M': 1000,
	}
	
	size, res := len(s), 0
	for i := 0; i < size; i++ {
		if i < size - 1 && m[s[i]] < m[s[i + 1]] {
			res -= m[s[i]]
		} else {
			res += m[s[i]]
		}
	}

	return res
}
// @lc code=end

