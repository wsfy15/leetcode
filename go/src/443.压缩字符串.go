/*
 * @lc app=leetcode.cn id=443 lang=golang
 *
 * [443] 压缩字符串
 */

// @lc code=start
func compress(chars []byte) int {
	i, j := 0, 0
	n := len(chars)
	for j < n {
		oldJ := j
		for j + 1 < n && chars[j] == chars[j + 1] {
			j++
		}

		chars[i] = chars[j]
		i++
		if j != oldJ {
			num := strconv.Itoa(j - oldJ + 1)
			for k := 0; k < len(num); k++ {
				chars[i] = num[k]
				i++
			}
		}

		j++
	}

	return i
}
// @lc code=end

