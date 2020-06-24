/*
 * @lc app=leetcode.cn id=345 lang=golang
 *
 * [345] 反转字符串中的元音字母
 */

// @lc code=start
func reverseVowels(s string) string {
	m := map[byte]bool{
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
	
	bytes := []byte(s)
	i, j := 0, len(bytes) - 1
	for i < j {
		for i < j {
			if _, ok := m[bytes[i]]; !ok {
				i++
			} else {
				break
			}
		}

		for i < j {
			if _, ok := m[bytes[j]]; !ok {
				j--
			} else {
				break
			}
		}

		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--
	}

	return string(bytes)
}
// @lc code=end

