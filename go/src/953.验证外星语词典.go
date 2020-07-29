/*
 * @lc app=leetcode.cn id=953 lang=golang
 *
 * [953] 验证外星语词典
 */

// @lc code=start
func isAlienSorted(words []string, order string) bool {
	m := map[byte]int{}
	for i := 0; i < len(order); i++ {
		m[order[i]] = i
	}

	for i := 1; i < len(words); i++ {
		n := min(len(words[i]), len(words[i-1]))
		j := 0
		for ; j < n; j++ {
			if m[words[i][j]] < m[words[i-1][j]] {
				return false
			} else if m[words[i][j]] > m[words[i-1][j]] {
				break
			}
		}

		if j == n { // 前n个都相同
			if len(words[i-1]) > n {
				return false
			}
		}
	}

	return true
}

func min(a ...int) int {
	res := a[0]
	for _, v := range a {
		if res > v {
			res = v
		}
	}
	return res
}
// @lc code=end

