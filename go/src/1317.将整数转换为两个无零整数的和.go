/*
 * @lc app=leetcode.cn id=1317 lang=golang
 *
 * [1317] 将整数转换为两个无零整数的和
 */

// @lc code=start
func getNoZeroIntegers(n int) []int {
	i, j := 1, n-1
	for i < j {
		if check(i) && check(j) {
			break
		}
		i++
		j--
	}
	return []int{i, j}
}

func check(n int) bool {
	for n > 0 {
		if n%10 == 0 {
			return false
		}
		n /= 10
	}
	return true
}

// @lc code=end

