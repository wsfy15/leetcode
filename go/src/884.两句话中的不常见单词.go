/*
 * @lc app=leetcode.cn id=884 lang=golang
 *
 * [884] 两句话中的不常见单词
 */

// @lc code=start
func uncommonFromSentences(A string, B string) []string {
	words := strings.Split(A, " ")
	counterA, counterB := map[string]int{}, map[string]int{}
	for _, word := range words {
		if word != "" {
			counterA[word]++
		}
	}

	words = strings.Split(B, " ")
	for _, word := range words {
		if word != "" {
			counterB[word]++
		}
	}

	res := []string{}
	for k, v := range counterA {
		if v == 1 && counterB[k] == 0 {
			res = append(res, k)
		}
	}
	for k, v := range counterB {
		if v == 1 && counterA[k] == 0 {
			res = append(res, k)
		}
	}

	return res
}
// @lc code=end

