/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	var cur []int
	used := make([]bool, len(nums))
	generate(nums, cur, &res, used, 0)
	return res
}

func generate(nums, cur []int, res *[][]int, used []bool, depth int) {
	if depth == len(nums) {
		// 需要拷贝到新数组，不能直接使用cur，否则可能因为调用函数对cur的修改而产生影响
		dst := make([]int, depth)
		copy(dst, cur)
		*res = append(*res, dst)
		// *res = append(*res, cur)
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}

		// 如果上一个相同数字已经用了，那下一个相同数字就可以产生新的排列
		// 如果上一个 没用，那当前这个和上一个 就是作用相同的，会产生重复
		if i > 0 && nums[i] == nums[i - 1] && !used[i - 1] {
			continue
		}

		used[i] = true
		generate(nums, append(cur, nums[i]), res, used, depth + 1)
		used[i] = false
	}
}
// @lc code=end

