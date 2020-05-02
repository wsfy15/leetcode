/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */

// @lc code=start
func combine(n int, k int) [][]int {
	res := [][]int{}
	dfs([]int{}, &res, 1, n, k)
	return res
}

func dfs(cur []int, res *[][]int, start, end, left int) {
	if left == 0 {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*res = append(*res, tmp)
		return
	}

	for i := start; i <= end - left + 1; i++ {
		tmp := append(cur, i)
		dfs(tmp, res, i + 1, end, left - 1)
	}
}

// @lc code=end

