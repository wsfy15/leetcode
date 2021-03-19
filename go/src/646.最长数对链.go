/*
 * @lc app=leetcode.cn id=646 lang=golang
 *
 * [646] 最长数对链
 */

// @lc code=start
func findLongestChain(pairs [][]int) int {
	res, n := 1, len(pairs)
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1] ||
			(pairs[i][1] == pairs[j][1] && pairs[i][0] < pairs[j][0])
	})

	pos := pairs[0][1]
	for i := 1; i < n; i++ {
		if pairs[i][0] > pos {
			res++
			pos = pairs[i][1]
		}
	} 

	return res
}
// @lc code=end

