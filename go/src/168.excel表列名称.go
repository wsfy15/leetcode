/*
 * @lc app=leetcode.cn id=168 lang=golang
 *
 * [168] Excel表列名称
 */

// @lc code=start
func convertToTitle(n int) string {
	sb := strings.Builder{}
	tmp := []byte{}
	for n > 0 {
		n--
		tmp = append(tmp, byte(n % 26) + 'A')
		n /= 26
	}

	for i := len(tmp) - 1; i >= 0; i-- {
		sb.WriteByte(tmp[i])
	}

	return sb.String()
}
// @lc code=end

