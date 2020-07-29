/*
 * @lc app=leetcode.cn id=917 lang=golang
 *
 * [917] 仅仅反转字母
 */

// @lc code=start
func reverseOnlyLetters(S string) string {
	left, right := 0, len(S) - 1
	if right < 1 {
		return S
	}

	bytes := []byte(S)

	for left < right {
		for left < right && !isAlpha(S[left]) {
			left++
		}
		for right > left && !isAlpha(S[right]) {
			right--
		}
		bytes[left], bytes[right] = bytes[right], bytes[left]
		right--
		left++
	}
	return string(bytes)
}


func isAlpha(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z')
}
// @lc code=end

