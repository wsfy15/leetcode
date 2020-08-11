/*
 * @lc app=leetcode.cn id=1108 lang=golang
 *
 * [1108] IP 地址无效化
 */

// @lc code=start
func defangIPaddr(address string) string {
	sb := strings.Builder{}
	for i := 0; i < len(address); i++ {
		if address[i] == '.' {
			sb.WriteString("[.]")
		} else {
			sb.WriteByte(address[i])
		}
	}

	return sb.String()
}

// @lc code=end

