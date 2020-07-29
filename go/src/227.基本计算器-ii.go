/*
 * @lc app=leetcode.cn id=227 lang=golang
 *
 * [227] 基本计算器 II
 */

// @lc code=start
func calculate(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	stack := []int{} // -1表示+、-3表示*、-4表示/
	sign := 1
	for i := 0; i < n; i++ {
		if s[i] == ' ' {
			continue
		}
		
		switch s[i] {
		case '+':
			stack = append(stack, -1)
		case '-':
			sign = -1
			stack = append(stack, -1)
		case '*':
			stack = append(stack, -3)
		case '/':
			stack = append(stack, -4)
		default:
			j := i
			for i < n && '0' <= s[i] && s[i] <='9' {
				i++
			}
			num, _ := strconv.Atoi(s[j:i])
			num *= sign
			sign = 1
			if len(stack) > 0 {
				if stack[len(stack) - 1] == -3 {
					num *= stack[len(stack) - 2]
					stack = stack[:len(stack) - 2]
				} else if stack[len(stack) - 1] == -4 {
					num = stack[len(stack) - 2] / num
					stack = stack[:len(stack) - 2]
				}
			}
			stack = append(stack, num)
			i--
		}
	}

	for len(stack) > 1 {
		num := 0
		switch stack[len(stack) - 2] {
		case -1:
			num = stack[len(stack) - 1] + stack[len(stack) - 3]
		case -3:
			num = stack[len(stack) - 3] * stack[len(stack) - 1]
		case -4:
			num = stack[len(stack) - 3] / stack[len(stack) - 1]
		}

		stack = stack[:len(stack) - 3]
		stack = append(stack, num)
	}

	return stack[0]
}
// @lc code=end

