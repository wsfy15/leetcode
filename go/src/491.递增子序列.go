/*
 * @lc app=leetcode.cn id=491 lang=golang
 *
 * [491] 递增子序列
 */

// @lc code=start
func findSubsequences(nums []int) [][]int {
	res := [][]int{}
	dfs(nums, []int{}, -101, 0, &res)

	return res
}

func dfs(nums []int, cur []int, lastNum, index int, res *[][]int) {
	if len(cur) >= 2 {
		dst := make([]int, len(cur))
		copy(dst, cur)
		*res = append(*res, dst)
	}

	m := map[int]bool{}
	for i := index; i < len(nums); i++ {
		if nums[i] < lastNum || m[nums[i]] {
			continue
		}
		m[nums[i]] = true
		cur = append(cur, nums[i])
		dfs(nums, cur, nums[i], i + 1, res)
		cur = cur[:len(cur) - 1]
	}
}
// @lc code=end

