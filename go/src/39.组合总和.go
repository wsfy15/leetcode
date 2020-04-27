/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	res, cur := [][]int{}, []int{}
	sort.Ints(candidates)
	generate(candidates, cur, &res, 0, target)

	return res
}

func generate(candidates, cur []int, res *[][]int, index, target int) {
	if target == 0 {
		// 这里必须拷贝创建一个新切片，否则用的仍是调用函数的cur，在调用函数中修改cur，会影响到res中的相应元素
		dst := make([]int, len(cur))
		copy(dst, cur)
		*res = append(*res, dst)
		return
	}

	if index == len(candidates) || target < 0 {
		return
	}

	generate(candidates, cur, res, index + 1, target)
	for target >= candidates[index] {
		target -= candidates[index]
		cur = append(cur, candidates[index])
		generate(candidates, cur, res, index + 1, target)
	}
}
// @lc code=end

