/*
 * @lc app=leetcode.cn id=1128 lang=golang
 *
 * [1128] 等价多米诺骨牌对的数量
 */

// @lc code=start
func numEquivDominoPairs(dominoes [][]int) int {
	counter := make([]int, 100)
	for _, domine := range dominoes {
		a, b := domine[0], domine[1]
		if a < b {
			counter[a*10+b]++
		} else {
			counter[b*10+a]++
		}
	}

	res := 0
	for i := 0; i < 100; i++ {
		if counter[i] > 1 {
			res += counter[i] * (counter[i] - 1) / 2
		}
	}
	return res
}

// @lc code=end

