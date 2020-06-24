/*
 * @lc app=leetcode.cn id=383 lang=golang
 *
 * [383] 赎金信
 */

// @lc code=start
func canConstruct(ransomNote string, magazine string) bool {
	letters := make([]int, 26)
	for i := 0; i < len(magazine); i++ {
		letters[magazine[i] - 'a']++
	}

	for i := 0; i < len(ransomNote); i++ {
		if letters[ransomNote[i] - 'a'] <= 0 {
			return false
		}
		letters[ransomNote[i] - 'a']--
	}

	return true
}
// @lc code=end

