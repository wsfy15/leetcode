/*
 * @lc app=leetcode.cn id=554 lang=golang
 *
 * [554] 砖墙
 */

// @lc code=start
func leastBricks(wall [][]int) int {
	counter := map[int]int{}
	n := len(wall)
	for i := 0; i < n; i++ {
		cur := 0
		for j := 0; j < len(wall[i])-1; j++ {
			cur += wall[i][j]
			counter[cur]++
		}
	}

	max := 0
	for _, v := range counter {
		if v > max {
			max = v
		}
	}

	return n - max
}

// @lc code=end

