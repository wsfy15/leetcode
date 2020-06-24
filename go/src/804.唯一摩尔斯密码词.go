/*
 * @lc app=leetcode.cn id=804 lang=golang
 *
 * [804] 唯一摩尔斯密码词
 */

// @lc code=start
func uniqueMorseRepresentations(words []string) int {
	m := []string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}
	res := 0
	exist := map[string]bool{}
	for _, word := range words {
		var str string
		for i := 0; i < len(word); i++ {
			str += m[word[i] - 'a']
		}
		if _, ok := exist[str]; !ok {
			res++
			exist[str] = true
		}
	}

	return res
}
// @lc code=end

