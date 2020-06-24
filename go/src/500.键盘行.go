/*
 * @lc app=leetcode.cn id=500 lang=golang
 *
 * [500] 键盘行
 */

// @lc code=start
func findWords(words []string) []string {
	res := []string{}
	lines := []string{"qwertyuiop", "asdfghjkl", "zxcvbnm"}
	m := map[byte]int{}
	for i, line := range lines {
		for j := 0; j < len(line); j++ {
			m[line[j]] = i
		}
	}

	for index, word := range words {
		word = strings.ToLower(word)
		line := m[word[0]]
		i := 1
		for ; i < len(word); i++ {
			if m[word[i]] != line {
				break
			}
		}

		if i == len(word) {
			res = append(res, words[index])
		}
	}

	return res
}
// @lc code=end

