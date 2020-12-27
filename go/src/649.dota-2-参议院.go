/*
 * @lc app=leetcode.cn id=649 lang=golang
 *
 * [649] Dota2 参议院
 */

// @lc code=start
func predictPartyVictory(senate string) string {
	n := len(senate)

	state := []byte(senate)
	r, d := true, true
	person := 0 // R可以淘汰人的次数

	for r && d {
		r, d = false, false
		for i := 0; i < n; i++ {
			if state[i] == 'R' {
				r = true
				if person < 0 {
					state[i] = '0'
				}
				person++
			} else if state[i] == 'D' {
				d = true
				if person > 0 {
					state[i] = '0'
				}
				person--
			}
		}
	}

	if person > 0 {
		return "Radiant"
	}
	return "Dire"
}

// @lc code=end

