/*
 * @lc app=leetcode.cn id=1160 lang=golang
 *
 * [1160] 拼写单词
 */

// @lc code=start
func countCharacters(words []string, chars string) int {
	var count = make(map[byte]int)
	for i := range chars {
			count[chars[i]]++
	}

	res := 0
	for i := range words {
			var tmp  = make(map[byte]int)
			for j := range words[i] {
					tmp[words[i][j]]++
			}
			flag := true
			for k, v := range tmp {
					if count[k] < v {
							flag = false
					}
			}
			if flag {
					res += len(words[i])
			}
	}
	return res
}
// @lc code=end

