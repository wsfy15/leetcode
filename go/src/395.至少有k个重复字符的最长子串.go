/*
 * @lc app=leetcode.cn id=395 lang=golang
 *
 * [395] 至少有K个重复字符的最长子串
 */

// @lc code=start
func longestSubstring(s string, k int) int {
	n := len(s)
	if n == 0 || k > n {
		return 0
	}
	if k <= 1 {
		return n
	}

	return dfs(s, 0, n - 1, k)
}

func dfs(s string, start, end, k int) int {
	if end - start + 1 < k {
		return 0
	}

	counter := [26]int{}
	for i := start; i <= end; i++ {
		counter[s[i] - 'a']++
	}

	for start < end && counter[s[start] - 'a'] < k {
		start++
	}

	for end > start && counter[s[end] - 'a'] < k {
		end--
	}

	if end - start + 1 < k {
		return 0
	}

	for i := start + 1; i < end; i++ {
		if counter[s[i] - 'a'] < k {
			return max(dfs(s, start, i - 1, k), dfs(s, i + 1, end, k))
		}
	}

	return end - start + 1
}

func max(a ...int) int {
	res := a[0]
	for _, v := range a {
		if res < v {
			res = v
		}
	}
	return res
}
// @lc code=end

