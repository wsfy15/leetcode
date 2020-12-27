/*
 * @lc app=leetcode.cn id=301 lang=golang
 *
 * [301] 删除无效的括号
 */

// @lc code=start
func removeInvalidParentheses(s string) []string {
	if len(s) == 0 {
		return []string{""}
	}

	var res []string
	queue := [][]byte{[]byte(s)}
	found := false
	for !found && len(queue) > 0 {
		newQueue := [][]byte{}
		exist := map[string]bool{}
		for _, str := range queue {
			if valid(str) {
				res = append(res, string(str))
				found = true
			}
			if found {
				continue
			}

			for i := 0; i < len(str); i++ {
				if str[i] != '(' && str[i] != ')' {
					continue
				}
				newStr := []byte{}
				newStr = append(newStr, str[:i]...)
				newStr = append(newStr, str[i+1:]...)
				if !exist[string((newStr))] {
					newQueue = append(newQueue, newStr)
					exist[string((newStr))] = true
				}
			}
		}
		queue = newQueue
	}

	return res
}

func valid(s []byte) bool {
	l := 0
	for _, ch := range s {
		if ch == '(' {
			l++
		} else if ch == ')' {
			l--
		}
		if l < 0 {
			return false
		}
	}

	return l == 0
}

// @lc code=end

