/*
 * @lc app=leetcode.cn id=30 lang=golang
 *
 * [30] 串联所有单词的子串
 */

// @lc code=start
func findSubstring(s string, words []string) []int {
	if len(s) == 0 || len(words) == 0 {
		return nil
	}

	wordLen := len(words[0])
	val, length := 0, len(words) * wordLen
	for _, v := range words {
		for i := 0; i < len(v); i++ {
			val += int(v[i])
		}
	}

	n, cur := len(s), 0
	res := []int{}
	if n < length {
		return nil
	}

	for i := 0; i < length - 1; i++ {
		cur += int(s[i])
	}

	for i := length - 1; i < n; i++ {
		cur += int(s[i])
		if cur == val {
			m := map[string]int{}
			for _, v := range words {
				m[v]++
			}

			j := i - length + 1
			for j <= i {
				if v, ok := m[s[j:j + wordLen]]; ok && v > 0 {
					m[s[j:j + wordLen]]--
					j = j + wordLen
				} else {
					break
				}

				if j == i + 1 {
					res = append(res, i - length + 1)
				}
			}
		}

		cur -= int(s[i - length + 1])
	}

	return res
}
// @lc code=end

