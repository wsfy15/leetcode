/*
 * @lc app=leetcode.cn id=748 lang=golang
 *
 * [748] 最短完整词
 */

// @lc code=start
func shortestCompletingWord(licensePlate string, words []string) string {
	licensePlate = strings.ToLower(licensePlate)
	m := make([]int, 26)
	for i := 0; i < len(licensePlate); i++ {
		if 'a' <= licensePlate[i] && licensePlate[i] <= 'z' {
			m[licensePlate[i] - 'a']++
		}
	}

	var res string
	for _, word := range words {
		m2 := make([]int, 26)
		copy(m2, m)
		for i := 0; i < len(word); i++ {
			m2[word[i] - 'a']--
		}

		i := 0
		for ; i < 26; i++ {
			if m2[i] > 0 {
				break
			}
		}
		if i == 26 {
			if len(res) == 0 || len(res) > len(word) {
				res = word
			}
		}
	}

	return res
}
// @lc code=end

