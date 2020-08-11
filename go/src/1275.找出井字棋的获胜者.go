/*
 * @lc app=leetcode.cn id=1275 lang=golang
 *
 * [1275] 找出井字棋的获胜者
 */

// @lc code=start
func tictactoe(moves [][]int) string {
	// 3列表示进制数的3位，3行表示二进制的三个数1、2、4
	pattern := map[int]bool{}
	pattern[0o007] = true
	pattern[0o070] = true
	pattern[0o700] = true
	pattern[0o111] = true
	pattern[0o222] = true
	pattern[0o444] = true
	pattern[0o421] = true
	pattern[0o124] = true

	a, b := 0, 0
	for i, move := range moves {
		val := (1 << move[0]) * (1 << (move[1] * 3))
		if i&1 == 1 {
			b |= val
			if _, ok := pattern[b]; ok {
				return "B"
			}
		} else {
			a |= val
			if _, ok := pattern[a]; ok {
				return "A"
			}
		}
	}

	if len(moves) == 9 {
		return "Draw"
	}
	return "Pending"
}

// @lc code=end

