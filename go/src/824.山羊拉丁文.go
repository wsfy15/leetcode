/*
 * @lc app=leetcode.cn id=824 lang=golang
 *
 * [824] 山羊拉丁文
 */

// @lc code=start
func toGoatLatin(S string) string {
	words := strings.Split(S, " ")
	sb := strings.Builder{}
	a := "a"

	vowel := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}

	for _, word := range words {
		if word == "" {
			continue
		}

		if vowel[word[0]] {
			sb.WriteString(word)
		} else {
			sb.WriteString(word[1:])
			sb.WriteByte(word[0])
		}
		sb.WriteString("ma")
		sb.WriteString(a)
		sb.WriteByte(' ')
		a += "a"
	}

	res := sb.String()
	if len(res) > 0 {
		return res[:len(res) - 1]
	}
	return ""
}
// @lc code=end

