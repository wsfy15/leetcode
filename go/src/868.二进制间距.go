/*
 * @lc app=leetcode.cn id=868 lang=golang
 *
 * [868] 二进制间距
 */

// @lc code=start
func binaryGap(N int) int {
	if N & (N - 1) == 0 {
		return 0
	}
	
	for N & 1 == 0 {
		N >>= 1
	}

	res, count := 0, 0
	for N > 0 {
		count++
		if N & 1 == 1 {
			res = max(res, count)
			count = 0
		}
		N >>= 1
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
// @lc code=end

