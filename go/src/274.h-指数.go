/*
 * @lc app=leetcode.cn id=274 lang=golang
 *
 * [274] H 指数
 */

// @lc code=start
func hIndex(citations []int) int {
	n := len(citations)
	sort.Ints(citations)
	for i := 0; i < n; i++ {
		// 当引用次数比剩余论文数量还多的时候，继续增大引用次数，只会减小论文数量，进而导致h变小
		if n-i <= citations[i] {
			return n - i
		}
	}

	return 0
}

// @lc code=end

