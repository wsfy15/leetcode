/*
 * @lc app=leetcode.cn id=90 lang=golang
 *
 * [90] 子集 II
 */

// @lc code=start
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	backtrack(nums, []int{}, 0, &res)
	return res
}

func backtrack(nums, cur []int, index int, res *[][]int) {
	dst := make([]int, len(cur))
	copy(dst, cur)
	*res = append(*res, dst)

	for i := index; i < len(nums); i++ {
		if i > index && nums[i] == nums[i - 1] {
			continue
		}

		cur = append(cur, nums[i])
		backtrack(nums, cur, i + 1, res)
		cur = cur[:len(cur) - 1]
	}
}

// @lc code=end

