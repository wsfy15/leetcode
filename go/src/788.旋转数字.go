/*
 * @lc app=leetcode.cn id=788 lang=golang
 *
 * [788] 旋转数字
 */

// @lc code=start
var m map[int]bool = map[int]bool {
	0: true, 1: true, 2: true, 5: true, 6: true, 8: true, 9: true,
}

var m2 map[int]bool = map[int]bool {
	2: true, 5: true, 6: true, 9: true,
}

func rotatedDigits(N int) int {
	res := 0
	for i := 2; i <= N; i++ {
		if isValid(i) {
			res++
		}
	}
	return res
}

func isValid(num int) bool {
	flag := false
	for num > 0 {
		digit := num % 10
		if _, ok := m[digit]; !ok {
			return false
		}
		if _, ok := m2[digit]; ok {
			flag = true
		}
		num /= 10
	}

	return flag
}
// @lc code=end

