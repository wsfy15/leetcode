/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] Z 字形变换
 */

// @lc code=start
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	n := len(s)
	sb := strings.Builder{}
	for i := 0; i < numRows; i++ {
		var delta []int
		if i == 0 || i == numRows - 1 {
			delta = []int{2 * numRows - 2}
		} else {
			delta = []int{2 * numRows - 2 * i - 2, 2 * i}
		}
		
		cur, j := 0, i
		for j < n {
			sb.WriteByte(s[j])
			j += delta[cur]
			cur = (cur + 1) % len(delta)
		}
	}

	return sb.String()
}
// @lc code=end

