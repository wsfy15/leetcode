/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
func subsets(nums []int) [][]int {
	n := len(nums)
	res := [][]int{{}}
	for i := 1; i <= n; i++ {
		dfs(nums, 0, i, []int{}, &res)
	}

	return res
}

func dfs(nums []int, start, count int, cur []int, res *[][]int) {
	if count == 0 {
		dst := make([]int, len(cur))
		copy(dst, cur)
		*res = append(*res, dst)
		return
	}

	for i := start; i < len(nums); i++ {
		tmp := append(cur, nums[i])
		dfs(nums, i + 1, count - 1, tmp, res)
	}
}

// @lc code=end

