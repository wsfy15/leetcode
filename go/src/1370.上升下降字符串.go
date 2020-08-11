/*
 * @lc app=leetcode.cn id=1370 lang=golang
 *
 * [1370] 上升下降字符串
 */

// @lc code=start
func sortString(s string) string {
	res := []byte{}
	count := [26]int{}
	for _, ch := range []byte(s) {
		count[ch-'a']++
	}

	for len(res) < len(s) {
		for i := 0; i < 26; i++ {
			if count[i] > 0 {
				count[i]--
				res = append(res, 'a'+byte(i))
			}
		}
		for i := 25; i >= 0; i-- {
			if count[i] > 0 {
				count[i]--
				res = append(res, 'a'+byte(i))
			}
		}
	}

	return string(res)
}

// @lc code=end

