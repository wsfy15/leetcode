/*
 * @lc app=leetcode.cn id=1078 lang=golang
 *
 * [1078] Bigram 分词
 */

// @lc code=start
func findOcurrences(text string, first string, second string) []string {
	words := strings.Split(text, " ")
	res := []string{}
	n := len(words)
	for i := 0; i < n - 2; i++ {
		if words[i] == first && words[i + 1] == second {
			res = append(res, words[i + 2])
		}
	}

	return res
}
// @lc code=end

