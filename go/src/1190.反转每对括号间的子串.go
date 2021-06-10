/*
 * @lc app=leetcode.cn id=1190 lang=golang
 *
 * [1190] 反转每对括号间的子串
 */

// @lc code=start
func reverseParentheses(s string) string {
	stack := []int{}
	res, cur := []byte{}, []byte(s)
	for i := 0; i < len(cur); i++ {
		if cur[i] == '(' {
			stack = append(stack, i)
		} else if cur[i] == ')' {
			left := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			reverse(cur, left+1, i-1)
		}
	}

	for _, ch := range cur {
		if ch == '(' || ch == ')' {
			continue
		}
		res = append(res, ch)
	}
	return string(res)
}

func reverse(cur []byte, left, right int) {
	for left < right {
		cur[left], cur[right] = cur[right], cur[left]
		left++
		right--
	}
}

// @lc code=end

