/*
 * @lc app=leetcode.cn id=925 lang=golang
 *
 * [925] 长按键入
 */

// @lc code=start
func isLongPressedName(name string, typed string) bool {
	n, m := len(name), len(typed)
	if n > m {
		return false
	}

	j := 0 // typed index
	for i := 0; i < n; i++ {
		if j < m && name[i] == typed[j] {
			j++
			if i + 1 < n && name[i] != name[i + 1] {
				for j < m && typed[j] == name[i] {
					j++
				}
			}
		} else {
			return false
		}
	}

	return j == m
}
// @lc code=end

