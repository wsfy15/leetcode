/*
 * @lc app=leetcode.cn id=224 lang=golang
 *
 * [224] 基本计算器
 */

// @lc code=start
import "fmt"

func calculate(s string) int {
	// 一个栈 存储正向扫描过程中的临时结果，当遇到(时入栈前面的结果和接下来结果的正负号，遇到)时弹栈
	var stack []int

	length := len(s)
	var res = 0
	var sign = 1

	var num int = 0 // 暂存读取中的数值
	for i := 0; i < length; i++ {
		switch s[i] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			num = num * 10 + int(s[i]) - int('0')

		case '(':
			stack = append(stack, res)
			stack = append(stack, sign)
			// 一切从新开始
			res = 0
			num = 0
			sign = 1

		case ')':
			res += sign * num
			res *= stack[len(stack) - 1] 
			res += stack[len(stack) - 2] 
			stack = stack[:len(stack) - 2] 
			num = 0

		case '+':
			res += sign * num
			sign = 1
			num = 0

		case '-':
			res += sign * num
			sign = -1
			num = 0

		default:
		}
	}

	if num != 0 {
		res += sign * num
	}
	return res
}
// @lc code=end

