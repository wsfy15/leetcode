/*
 * @lc app=leetcode.cn id=937 lang=golang
 *
 * [937] 重新排列日志文件
 */

// @lc code=start
func reorderLogFiles(logs []string) []string {
	n, index := len(logs)-1, len(logs)-1
	for n >= 0 {
		parts := strings.Split(logs[n], " ")
		if isDigit(parts[1][0]) {
			logs[n], logs[index] = logs[index], logs[n]
			index--
		}
		n--
	}

	sort.Slice(logs[:index+1], func(i, j int) bool {
		iparts, jparts := strings.Split(logs[i], " "), strings.Split(logs[j], " ")
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

