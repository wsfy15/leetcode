/*
 * @lc app=leetcode.cn id=1374 lang=golang
 *
 * [1374] 生成每种字符都是奇数个的字符串
 */

// @lc code=start
func generateTheString(n int) string {
	bytes := make([]byte, n)
	for i := 0; i < n; i++ {
		bytes[i] = 'a'
	}
	if n&1 == 0 {
		bytes[n-1] = 'b'
	}

	return string(bytes)
}

// @lc code=end

