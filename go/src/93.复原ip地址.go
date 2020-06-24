/*
 * @lc app=leetcode.cn id=93 lang=golang
 *
 * [93] 复原IP地址
 */

// @lc code=start
func restoreIpAddresses(s string) []string {
	res := []string{}
	if len(s) < 4 {
		return nil
	}

	backtrace(s, 0, 0, "", &res)
	return res
}

func backtrace(s string, index, count int, cur string, res *[]string) {
	if count == 4 {
		if index == len(s) {
			*res = append(*res, cur[:len(cur) - 1])
		}
		return
	}

	if index >= len(s) {
		return
	}

	if s[index] == '0' {
		backtrace(s, index + 1, count + 1, cur + "0.", res)
		return
	}

	for i := index; i < len(s) && i < index + 3; i++ {
		if v, _ := strconv.Atoi(s[index: i + 1]); v < 256 {
			backtrace(s, i + 1, count + 1, cur + s[index: i+1] + ".", res)
		} else {
			break
		}
	}
}

// @lc code=end

