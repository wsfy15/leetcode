/*
 * @lc app=leetcode.cn id=781 lang=golang
 *
 * [781] 森林中的兔子
 */

// @lc code=start
func numRabbits(answers []int) int {
	res, n := 0, len(answers)
	sort.Ints(answers)

	// 设有某种颜色的兔子 m 只，它们回答的答案数值为 cnt 那么m=cnt+1
	// 如果同个数值的出现次数大于cnt+1，说明有一种新颜色出现
	for i := 0; i < n; i++ {
		res += answers[i] + 1
		k := answers[i]
		for i+1 < n && k > 0 && answers[i+1] == answers[i] {
			i++
			k--
		}
	}

	return res
}

// @lc code=end

