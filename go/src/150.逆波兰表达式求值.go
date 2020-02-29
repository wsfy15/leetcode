/*
 * @lc app=leetcode.cn id=150 lang=golang
 *
 * [150] 逆波兰表达式求值
 */

// @lc code=start
func evalRPN(tokens []string) int {
		n := len(tokens)
		if n == 0 {
			return 0
		}

		stack := []int{}
		for i := 0; i < n; i++ {
			if num, err:=strconv.Atoi(tokens[i]); err != nil {
				former := stack[len(stack) - 2]
				later := stack[len(stack) - 1]
				stack = stack[:len(stack) - 2]
				res := 0
				switch tokens[i] {
				case "+":
					res = former + later
				case "-":
					res = former - later
				case "*":
					res = former * later
				case "/":
					res = former / later
				}

				stack = append(stack, res)
			} else {
				stack = append(stack, num)
			}
		}

		return stack[0]
}
// @lc code=end

