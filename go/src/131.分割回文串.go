/*
 * @lc app=leetcode.cn id=131 lang=golang
 *
 * [131] 分割回文串
 */

// @lc code=start
func partition(s string) [][]string {
	res := [][]string{}
	backtrace(s, 0, []string{}, &res)
	return res
}

func backtrace(s string, start int, cur []string, res *[][]string) {
	if start == len(s) {
		dst := make([]string, len(cur))
		copy(dst, cur)
		*res = append(*res, dst)
		return
	}

	for i := start; i < len(s); i++ {
		if isParli(s[start:i+1]) {
			cur = append(cur, s[start:i+1])
			backtrace(s, i + 1, cur, res)
			cur = cur[:len(cur) - 1]
		}
	}
}

func isParli(s string) bool {
	start, end := 0, len(s) - 1
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}

// @lc code=end

