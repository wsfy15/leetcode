/*
 * @lc app=leetcode.cn id=38 lang=golang
 *
 * [38] 外观数列
 */

// @lc code=start
func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	s := countAndSay(n - 1)
	sb := strings.Builder{}
	index, size := 1, len(s)
	var count byte = 1
	var last byte = s[0]
	for index < size {
		if last == s[index] {
			count++
		} else {
			sb.WriteByte(count + '0')
			sb.WriteByte(last)
			count = 1
			last = s[index]
		}
		index++
	}

	sb.WriteByte(count + '0')
	sb.WriteByte(last)

	return sb.String()
}
// @lc code=end

