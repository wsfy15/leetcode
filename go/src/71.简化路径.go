/*
 * @lc app=leetcode.cn id=71 lang=golang
 *
 * [71] 简化路径
 */

// @lc code=start
func simplifyPath(path string) string {
	stack := []string{}
	n := len(path)
	for i := 0; i < n; {
		j := i + 1
		for j < n && path[j] != '/' {
			j++
		}

		str := path[i+1:j]
		if str == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack) - 1]
			}
		} else if str != "." && str != "" {
			stack = append(stack, str)
		}
		i = j
	}

	str := strings.Join(stack,  "/")
	return "/" + str
}
// @lc code=end

