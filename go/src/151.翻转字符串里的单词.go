/*
 * @lc app=leetcode.cn id=151 lang=golang
 *
 * [151] 翻转字符串里的单词
 */

// @lc code=start
func reverseWords(s string) string {
	size := len(s)
	i := 0
	for i < size && s[i] == ' ' {
		i++
	} 

	var strs []string
	for ; i < size; {
		j := i + 1
		for j < size && s[j] != ' ' {
			j++
		}
		strs = append(strs, s[i:j])
		for j < size && s[j] == ' ' {
			j++
		}
		i = j
	}
	
	var res string
	for i := len(strs) - 1; i >= 0; i-- {
		res += strs[i] 
		res += " "
	}
	
	if len(res) == 0 {
		return ""
	}
	return res[0: len(res) - 1]
}
// @lc code=end

