/*
 * @lc app=leetcode.cn id=995 lang=golang
 *
 * [995] K 连续位的最小翻转次数
 */

// @lc code=start
func minKBitFlips(A []int, K int) int {
	n, res := len(A), 0
	queue := []int{}
	for i := 0; i < n; i++ {
		for len(queue) > 0 && queue[0] + K <= i {
			queue = queue[1:]
		}

		if A[i] == len(queue) % 2 {
			if i + K > n {
				return -1
			}
			queue = append(queue, i)
			res++
		}
	}

	return res
}
// @lc code=end

