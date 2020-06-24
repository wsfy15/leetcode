/*
 * @lc app=leetcode.cn id=771 lang=golang
 *
 * [771] 宝石与石头
 */

// @lc code=start
func numJewelsInStones(J string, S string) int {
	m := map[byte]bool{}
	for i := 0; i < len(J); i++ {
		m[J[i]] = true
	}

	res := 0
	for i := 0; i < len(S); i++ {
		if _, ok := m[S[i]]; ok {
			res++
		}
	}
	return res
}
// @lc code=end

