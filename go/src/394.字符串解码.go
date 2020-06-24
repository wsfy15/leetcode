/*
 * @lc app=leetcode.cn id=394 lang=golang
 *
 * [394] 字符串解码
 */

// @lc code=start
type node struct {
	typ int // 0: [ 1: number 2: string
	val string
}


func decodeString(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}

	stack := []node{}
	for index := 0; index < n; index++ {
		switch s[index] {
		case '[':
			stack = append(stack, node{ typ: 0 })
		case ']':
			str := stack[len(stack) - 1].val
			stack = stack[:len(stack) - 2] // pop string and [

			count := 1
			if stack[len(stack) - 1].typ == 1 {
				count, _ = strconv.Atoi(stack[len(stack) - 1].val)
				stack = stack[:len(stack) - 1] // pop number
			}
			str = strings.Repeat(str, count)

			if len(stack) > 0 && stack[len(stack) - 1].typ == 2 {
				stack[len(stack) - 1].val += str
			} else {
				stack = append(stack, node{ typ: 2, val: str })
			}

		default:
			if '0' <= s[index] && s[index] <= '9' {
				start := index
				for index < n && '0' <= s[index] && s[index] <= '9' {
					index++
				}
				stack = append(stack, node{ typ: 1, val: s[start: index] })
			} else {
				start := index
				for index < n && s[index] != ']' && (s[index] < '0' || s[index] > '9') {
					index++
				}
				if len(stack) > 0 && stack[len(stack) - 1].typ == 2 {
					stack[len(stack) - 1].val += s[start: index]
				} else {
					stack = append(stack, node{ typ: 2, val: s[start: index] })
				}
			}
			index--
		}
	}

	for len(stack) > 1 {
		stack[len(stack) - 2].val += stack[len(stack) - 1].val
		stack = stack[:len(stack) - 1]
	}

	return stack[0].val
}
// @lc code=end

