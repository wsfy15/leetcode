/*
 * @lc app=leetcode.cn id=937 lang=golang
 *
 * [937] 重新排列日志文件
 */

// @lc code=start
func reorderLogFiles(logs []string) []string {
	sort.Slice(logs, func(i, j int) bool {
		iparts, jparts := strings.Split(logs[i], " "), strings.Split(logs[j], " ")
		if isDigit(iparts[1][0]) && isDigit(jparts[1][0]) {
			return i < j
		} else if isDigit(iparts[1][0]) {
			return false
		} else if isDigit(jparts[1][0]) {
			return true
		}

		iid, jid := iparts[0], jparts[0]
		if logs[i][len(iid):] == logs[j][len(jid):] {
			return iid < jid
		}
		return logs[i][len(iid):] < logs[j][len(jid):]
	})

	return logs
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
// @lc code=end

