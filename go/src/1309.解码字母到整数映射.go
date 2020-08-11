/*
 * @lc app=leetcode.cn id=1309 lang=golang
 *
 * [1309] 解码字母到整数映射
 */

// @lc code=start
func freqAlphabets(s string) string {
	words := strings.Split(s, "#")
	bytes := []byte{}
	n := len(words)
	for i := 0; i < n-1; i++ {
		word := words[i]
		if len(word) > 2 {
			for i := 0; i < len(word)-2; i++ {
				bytes = append(bytes, 'a'+word[i]-'1')
			}
		}
		num, _ := strconv.Atoi(word[len(word)-2 : len(word)])
		bytes = append(bytes, 'a'+byte(num)-1)
	}

	for i := 0; i < len(words[n-1]); i++ {
		bytes = append(bytes, 'a'+words[n-1][i]-'1')
	}

	return string(bytes)
}

// @lc code=end

