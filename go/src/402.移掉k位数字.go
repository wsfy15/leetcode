/*
 * @lc app=leetcode.cn id=402 lang=golang
 *
 * [402] 移掉K位数字
 */

// @lc code=start
func removeKdigits(num string, k int) string {
	n := len(num)
	if k >= n {
		return "0"
	}

	stack := []byte{}
	for i := 0; i < n; i++ {
		for len(stack) > 0 && k > 0 && num[i] < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, num[i])
	}

	if k > 0 {
		stack = stack[:len(stack)-k]
	}

	for len(stack) > 0 && stack[0] == '0' {
		stack = stack[1:]
	}

	if len(stack) == 0 {
		return "0"
	}

	return string(stack)
}

// @lc code=end

