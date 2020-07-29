/*
 * @lc app=leetcode.cn id=997 lang=golang
 *
 * [997] 找到小镇的法官
 */

// @lc code=start
func findJudge(N int, trust [][]int) int {
	m := map[int]bool{}
	counter := map[int]int{}
	for _, item := range trust {
		m[item[0]] = true
		counter[item[1]]++
	}

	for i := 1; i <= N; i++ {
		if !m[i] && counter[i] == N - 1 {
			return i
		}
	}
	return -1
}
// @lc code=end

