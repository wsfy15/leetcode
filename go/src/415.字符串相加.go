/*
 * @lc app=leetcode.cn id=415 lang=golang
 *
 * [415] 字符串相加
 */

// @lc code=start
func addStrings(num1 string, num2 string) string {
	res := []byte{}
	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}

	n, m := len(num1), len(num2)
	var carry byte = 0
	for i := 0; i < n; i++ {
		var tmp byte = num1[i] - '0' + carry
		carry = 0
		if i < m {
			tmp += num2[i] - '0'
		}
		if tmp > 9 {
			carry = 1
			tmp -= 10
		}
		res = append(res, tmp + '0')
	}

	if carry > 0 {
		res = append(res, carry)
	}

	i, j := 0, len(res) - 1
	for i < j {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}
// @lc code=end

