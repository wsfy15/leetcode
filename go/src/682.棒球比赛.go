/*
 * @lc app=leetcode.cn id=682 lang=golang
 *
 * [682] 棒球比赛
 */

// @lc code=start
import "strconv"

func calPoints(ops []string) int {
	// 用一个栈存储过程中的分数
	var stack []int
	var res int = 0
	for i := 0; i < len(ops); i++ {
		switch ops[i] {
		case "C":
			stack = stack[:len(stack) - 1]

		case "D":
			stack = append(stack, 2 * stack[len(stack) - 1])

		case "+":
			stack = append(stack, stack[len(stack) - 1] + stack[len(stack) - 2])

		default:
			val, _ := strconv.Atoi(ops[i])
			stack = append(stack, val)
		}
	}


	for i := 0; i < len(stack); i++ {
		res += stack[i]
	}
	
	return res
}
// @lc code=end

