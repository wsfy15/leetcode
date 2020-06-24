/*
 * @lc app=leetcode.cn id=482 lang=golang
 *
 * [482] 密钥格式化
 */

// @lc code=start
func licenseKeyFormatting(S string, K int) string {
	words := strings.Split(S, "-")
	n := 0
	for _, word := range words {
		n += len(word)
	}
	newWords := []string{}
	var newWord []byte
	left := n % K
	for _, word := range words {
		for i := 0; i < len(word); i++ {
			if left == 0 {
				if len(newWord) > 0 {
					newWords = append(newWords, string(newWord))
					newWord = []byte{}
				}
				left = K
			}

			if 'a' <= word[i] && word[i] <= 'z' {									
				newWord = append(newWord, word[i] - 32)
			} else {
				newWord = append(newWord, word[i])
			}
			left--
		}
	}
	newWords = append(newWords, string(newWord))
	return strings.Join(newWords, "-")
}
// @lc code=end

