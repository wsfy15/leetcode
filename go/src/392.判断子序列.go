/*
 * @lc app=leetcode.cn id=392 lang=golang
 *
 * [392] 判断子序列
 */

// @lc code=start
func isSubsequence(s string, t string) bool {
	// 后续挑战
	// 存储每个字母在t中的下标
	counter := make([][]int, 26)
	for i := 0; i < len(t); i++ {
		counter[t[i] - 'a'] = append(counter[t[i] - 'a'], i)
	}

	index := -1
	for i := 0; i < len(s); i++ {
		letter := s[i] - 'a'
		if len(counter[letter]) == 0 {
			return false
		}

		left, right := 0, len(counter[letter]) - 1
		for left < right {
			mid := left + (right - left) >> 1
			if counter[letter][mid] <= index {
				left = mid + 1
			} else {
				right = mid
			}
		}
		if counter[letter][left] <= index {
			return false
		}
		index = counter[letter][left]
	}

	return true
}

func isSubsequence2(s string, t string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == len(s)
}
// @lc code=end

