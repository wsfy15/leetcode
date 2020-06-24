/*
 * @lc app=leetcode.cn id=290 lang=golang
 *
 * [290] 单词规律
 */

// @lc code=start
func wordPattern(pattern string, str string) bool {
	m := len(pattern)
	recorder := map[byte]string{}
	used := map[string]bool{}
	words := strings.FieldsFunc(str, func(r rune) bool {
		return r == ' '
	})

	if len(words) != m {
		return false
	}

	for i := 0; i < m; i++ {
		if v, ok := recorder[pattern[i]]; ok {
			if v != words[i] {
				return false
			}
		} else {
			if !used[words[i]] {
				recorder[pattern[i]] = words[i]
				used[words[i]] = true
			} else {
				return false
			}
		}
	}

	return true
}
// @lc code=end

