/*
 * @lc app=leetcode.cn id=205 lang=golang
 *
 * [205] 同构字符串
 */

// @lc code=start
func isIsomorphic(s string, t string) bool {
	n := len(s)
	trans := map[byte]byte{}
	used := map[byte]bool{}

	for i := 0; i < n; i++ {
		if v, ok := trans[s[i]]; ok {
			if v != t[i] {
				return false
			}
		} else {
			if used[t[i]] {
				return false
			}
			trans[s[i]] = t[i]
			used[t[i]] = true
		}
	}

	return true
}
// @lc code=end

