/*
 * @lc app=leetcode.cn id=216 lang=golang
 *
 * [216] 组合总和 III
 */

// @lc code=start
func combinationSum3(k int, n int) [][]int {
	var res [][]int
	backtrack([]int{}, k, n, 0, 1, &res)
	return res
}

func backtrack(cur []int, k, n, total, num int, res *[][]int) {
	if total == n && len(cur) == k {
		dst := make([]int, k)
		copy(dst, cur)
		*res = append(*res, dst)
		return
	}

	if num > 9 || len(cur) > k || (k - len(cur)) * num + total > n {
		return
	}

	for i := num; i <= 9; i++ {
		if total + i <= n {
			cur = append(cur, i)
			backtrack(cur, k, n, total + i, i + 1, res)
			cur = cur[:len(cur) - 1]
		} else {
			break
		}
	}
}
// @lc code=end

