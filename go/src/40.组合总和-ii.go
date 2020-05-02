/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 */

// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res, cur := [][]int{}, []int{}
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

	// 在每一层，都可以选择[index, len(candidates)) 任一元素
	for i := index; i < len(candidates); i++ {
		if i > index && candidates[i] == candidates[i - 1] { // 在当前层 出现重复
			continue
		}
		if target < candidates[i] {
			break
		}

		generate(candidates, append(cur, candidates[i]), res, i + 1, target - candidates[i])
	}
}
// @lc code=end

