/*
 * @lc app=leetcode.cn id=1178 lang=golang
 *
 * [1178] 猜字谜
 */

// @lc code=start
func findNumOfValidWords(words []string, puzzles []string) []int {
	n, m := len(words), len(puzzles)
	counter := map[int]int{} // 每种symbol的数量
	for i := 0; i < n; i++ {
		symbol := 0
		for _, ch := range words[i] {
			symbol |= 1 << (ch - 'a') 
		}
		counter[symbol]++
	}

	res := make([]int, m)
	for i := 0; i < m; i++ {
		symbol := 0
		for _, ch := range puzzles[i] {
			symbol |= 1 << (ch - 'a')
		}
		for j := symbol; j > 0; j = (j - 1) & symbol {
			if 1 << (puzzles[i][0] - 'a') & j > 0 {
				res[i] += counter[j]
			}
		}
	}

	return res
}
// @lc code=end

