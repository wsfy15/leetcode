/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) [][]int {
	return generate(nums, 0)
}

func generate(nums []int, index int) [][]int {
	if index == len(nums) {
		return [][]int{{}}
	}

	var res [][]int
	for i := index; i < len(nums); i++ {
		nums[index], nums[i] = nums[i], nums[index]
		tmp := generate(nums, index + 1)
		for _, v := range tmp {
			res = append(res, append(v, nums[index]))
		}
		nums[index], nums[i] = nums[i], nums[index]
	}

	return res
}
// @lc code=end

