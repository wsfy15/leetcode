/*
 * @lc app=leetcode.cn id=65 lang=golang
 *
 * [65] 有效数字
 */

// @lc code=start
func isNumber(s string) bool {
	table := [][]int{
		{ 0, 1, 6, 2,-1,-1},
		{-1,-1, 6, 2,-1,-1},
		{-1,-1, 3,-1,-1,-1},
		{ 8,-1, 3,-1, 4,-1},
		{-1, 7, 5,-1,-1,-1},
		{ 8,-1, 5,-1,-1,-1},
		{ 8,-1, 6, 3, 4,-1},
		{-1,-1, 5,-1,-1,-1},
		{ 8,-1,-1,-1,-1,-1},
	}
	finals := []byte{0,0,0,1,0,1,1,0,1}
	
	state := 0
	for i := 0; i < len(s); i++ {
		col := 0
		switch s[i] {
		case 'e':
			col = 4
		case ' ':
			col = 0
		case '+', '-':
			col = 1
		case '.':
			col = 3
		default:
			if '0' <= s[i] && s[i] <= '9' {
				col = 2
			} else {
				col = 5
			}
		}
		state = table[state][col]
		if state == -1 {
			return false
		}
	}

	return finals[state] == 1
}
// @lc code=end

