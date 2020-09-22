/*
 * @lc app=leetcode.cn id=214 lang=golang
 *
 * [214] 最短回文串
 */

// @lc code=start
func shortestPalindrome(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}

	// s中包含的回文串必须包括左端点
	rev := Reverse(s)
	pos := -1
	pow, b := 1, 26
	hash1, hash2 := 0, 0 // 正序和逆序的哈希值
	for i := 0; i < n; i++ {
		hash1 = hash1*b + int(s[i]-'a') + 1
		hash2 = hash2 + int(s[i]-'a'+1)*pow
		if hash1 == hash2 {
			if s[:i+1] == rev[n-i-1:] {
				pos = i
			}
		}
		pow *= b
	}

	return rev[:n-pos-1] + s
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// @lc code=end

