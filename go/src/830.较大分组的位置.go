/*
 * @lc app=leetcode.cn id=830 lang=golang
 *
 * [830] 较大分组的位置
 */

// @lc code=start
func largeGroupPositions(S string) [][]int {
	n := len(S)
	if n < 3 {
		return nil
	}

	res := [][]int{}
	for i := 0; i < n; i++ {
		j := i + 1
		for j < n && S[i] == S[j] {
			j++
		}

		if j - i >= 3 {
			res = append(res, []int{i, j - 1})
		}
		i = j - 1
	}

	return res
}
// @lc code=end

