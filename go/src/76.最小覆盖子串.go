/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 */

// @lc code=start
func minWindow(s string, t string) string {
	n, m := len(s), len(t)
	if n < m {
		return ""
	}

	counter := map[byte]int{}
	for i := 0; i < m; i++ {
		counter[t[i]]++
	}

	count := 0
	left, right := 0, 0
	var res string
	for right < n {
		if v, ok := counter[s[right]]; ok {
			counter[s[right]]--

			if v > 0 {
				count++
				for count == m && left < n {
					if len(res) == 0 || len(res) > right - left + 1 {
						res = s[left:right+1]
					}

					if _, ok := counter[s[left]]; ok {
						counter[s[left]]++
						if counter[s[left]] > 0 {
							count--
						}
					}
					left++
				}
			}
		}
		right++
	}

	return res
}
// @lc code=end

