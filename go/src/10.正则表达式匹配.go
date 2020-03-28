/*
 * @lc app=leetcode.cn id=10 lang=golang
 *
 * [10] 正则表达式匹配
 */

// @lc code=start
var states [][]int
func isMatch(s string, p string) bool {
	sl := len(s)
	pl := len(p)
	if pl == 0 {
		return sl == 0
	}

	// states[i][j] 表示p[0...i]与s[0...j] 匹配是否成功 0: 未计算 1：匹配 2：不匹配
	states = make([][]int, pl + 1)
	for i := range states {
		states[i] = make([]int, sl + 1)
	}

	return dp(0, 0, pl, sl, s, p)
} 

// i: pattern index
// j: string index
func dp(i, j, pl, sl int, s, p string) bool {
	if states[i][j] > 0 {
		return states[i][j] == 1
	}

	if i == pl {
		states[i][j] = 1
		if j != sl {
			states[i][j] = 2
		}
		return states[i][j] == 1
	}

	if j >= sl {
		for i < pl - 1 {
			if p[i+1] == '*' {
				i += 2
			} else {
				return false
			}
		}
		return i == pl
	}

	if p[i] == s[j] || p[i] == '.' {
		if i < pl - 1 && p[i+1] == '*' {
			// 匹配1位主串当前字符并保持* || 匹配1位主串当前字符，不保持* || 匹配0位
			if dp(i, j+1, pl, sl, s, p) || dp(i+2, j+1, pl, sl, s, p) || dp(i+2, j, pl, sl, s, p) {
				states[i][j] = 1
			} else {
				states[i][j] = 2
			}
		} else {
			if dp(i+1, j+1, pl, sl, s, p) {
				states[i][j]  = 1
			} else {
				states[i][j] = 2 
			}
		}
	} else if i < pl - 1 && p[i+1] == '*' { // 虽然没匹配上，但如果模式串下一位是*，就可以匹配主串0位字符
		if dp(i+2, j, pl, sl, s, p) {
			states[i][j]  = 1
		} else {
			states[i][j] = 2 
		}
	}
	return states[i][j] == 1
}
// @lc code=end

