/*
 * @lc app=leetcode.cn id=1111 lang=golang
 *
 * [1111] 有效括号的嵌套深度
 */

// @lc code=start
func maxDepthAfterSplit(seq string) []int {
	size := len(seq)
	res := make([]int, size)

	for i := 0; i < size; i++ {
		if seq[i] == '(' {
			res[i] = i & 1
		} else {
			res[i] = (i + 1) & 1
		}
	}

	return res
}
// @lc code=end

