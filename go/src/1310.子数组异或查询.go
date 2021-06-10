/*
 * @lc app=leetcode.cn id=1310 lang=golang
 *
 * [1310] 子数组异或查询
 */

// @lc code=start
func xorQueries(arr []int, queries [][]int) []int {
	n, m, cur := len(arr), len(queries), 0
	prefix := make([]int, n+1)
	for i := 0; i < n; i++ {
		cur ^= arr[i]
		prefix[i+1] = cur
	}

	res := make([]int, m)
	for i := 0; i < m; i++ {
		l, r := queries[i][0], queries[i][1]+1
		res[i] = prefix[l] ^ prefix[r]
	}

	return res
}

// @lc code=end

