/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */

// @lc code=start
import "math"


func myAtoi(str string) int {
	size := len(str)
	if size == 0 {
		return 0
	}

	var res int64
	i := 0
	for i < size && str[i] == ' ' {
		i++
	}

	if i == size || ((str[i] != '-' && str[i] != '+') && (str[i] > '9' || str[i] < '0')) {
		return 0
	}

	var sign int64 = 1
	if str[i] == '-' {
		sign = -1
		i++
	} else if str[i] == '+' {
		i++
	}

	for i < size && str[i] >= '0' && str[i] <= '9' {
		res = res*10 + sign * int64(str[i] - '0')

		if res > math.MaxInt32  {
			return math.MaxInt32
		} else if (res < math.MinInt32) {
			return math.MinInt32
		}

		i++
	}

	return int(res)
}
// @lc code=end

