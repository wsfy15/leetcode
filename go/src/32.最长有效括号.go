/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] 最长有效括号
 */

// @lc code=start
func longestValidParentheses(s string) int {
	// 对于遇到的每个 ( ，将它的下标放入栈中
	// 对于遇到的每个 ) ，弹出栈顶的元素，
	//     如果栈为空，说明不是有效的，所以将当前元素的下标压栈，该下标有可能是下一个有效括号字符串的下标 的前一位 
	//     栈不为空，将当前元素的下标与当前栈顶元素下标作差，得出当前有效括号字符串的长度

	size := len(s)
	max := 0
	var stack []int
	stack = append(stack, -1)	// 先放入-1是为了 [(, )] 这种清空

	for i := 0; i < size; i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack) - 1]
			if len(stack) == 0 {
				stack = append(stack, i)	// 下一个有效括号字符串的下标 的前一位 可能就是当前)的位置
			} else {
				tmp := i - stack[len(stack) - 1]
				if tmp > max {
					max = tmp
				}
			}
		}
	}

	return max
}

func longestValidParentheses2(s string) int {
	// 整个括号串中，要么只多余左括号，要么只多余右括号，要么多余的右括号必定在多余的左括号的左侧

	size := len(s)
	max := 0
	leftBraceCount := 0 // ( 的个数
	rightBraceCount := 0

	// 从左向右遍历时，相当于按多余的右括号来划分区间，此时无需考虑多余左括号的情况[(, ), (, (, (, )]，这种情况通过从右向左遍历可以解决
  for i := 0; i < size; i++ {
		if s[i] ==	'(' {
			leftBraceCount++
		} else {
			rightBraceCount++
		}

		if rightBraceCount > leftBraceCount {
			rightBraceCount = 0
			leftBraceCount = 0
		} else if rightBraceCount == leftBraceCount {
			if rightBraceCount * 2 > max {
				max = rightBraceCount * 2
			}
		}
	}

	// 从右向左遍历时，相当于按多余的左括号来划分区间
	leftBraceCount = 0
	rightBraceCount = 0
	for i := size - 1; i >= 0; i-- {
		if s[i] ==	'(' {
			leftBraceCount++
		} else {
			rightBraceCount++
		}

		if leftBraceCount > rightBraceCount {
			rightBraceCount = 0
			leftBraceCount = 0
		} else if rightBraceCount == leftBraceCount {
			if rightBraceCount * 2 > max {
				max = rightBraceCount * 2
			}
		}
	}


	return max
}
// @lc code=end

