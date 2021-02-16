/*
 * @lc app=leetcode.cn id=424 lang=golang
 *
 * [424] 替换后的最长重复字符
 */

// @lc code=start
func characterReplacement(s string, k int) int {
	n := len(s)
	if k >= n-1 {
		return n
	}

	counter := make([]int, 26)
	curMax, res := 0, 0
	left, right := 0, 0
	for right < n {
		counter[s[right] - 'A']++
		if counter[s[right] - 'A'] > curMax {
			curMax = counter[s[right] - 'A']
		}

		for right - left + 1 > curMax + k {
			counter[s[left] - 'A']--
			left++
			curMax = max(counter...)
		}
		
		if curMax + k > res {
			res = min(curMax + k, n)
		}

		right++
	}

	return res
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

func min(a ...int) int {
	res := a[0]
	for _, v := range a {
		if res > v {
			res = v
		}
	}
	return res
}

// @lc code=end

