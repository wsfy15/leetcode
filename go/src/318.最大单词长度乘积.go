/*
 * @lc app=leetcode.cn id=318 lang=golang
 *
 * [318] 最大单词长度乘积
 */

// @lc code=start
func maxProduct(words []string) int {
	n := len(words)
	bits := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < len(words[i]); j++ {
			bits[i] |= 1 << (words[i][j] - 'a')
		}
	}

	res := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if bits[i] & bits[j] == 0 {
				res = max(res, len(words[i]) * len(words[j]))
			}
		}
	}

	return res
}

func max(a ...int) int {
	ans := a[0]
	for _, v := range a {
		if v > ans {
			ans = v
		}
	}
	return ans
}
// @lc code=end

