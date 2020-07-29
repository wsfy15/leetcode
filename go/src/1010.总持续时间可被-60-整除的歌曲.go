/*
 * @lc app=leetcode.cn id=1010 lang=golang
 *
 * [1010] 总持续时间可被 60 整除的歌曲
 */

// @lc code=start
func numPairsDivisibleBy60(time []int) int {
	res := 0
	count := [60]int{}
	for _, t := range time {
		t %= 60
		res += count[60 - t]
		count[t]++
	}
	return res
}
// @lc code=end

