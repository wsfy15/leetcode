/*
 * @lc app=leetcode.cn id=994 lang=golang
 *
 * [994] 腐烂的橘子
 */

// @lc code=start
// 采用队列会更高效，用一个哈希表记录每个位置的橘子被感染的时间 
// 初始时队列里只有已经感染的橘子的位置（row*numPerRow + col)
// 然后出队，将新感染的入队，同时记录该橘子被感染的时间 出队的橘子因为已经把其附近的感染了，所以不用再入队
// 记录新鲜橘子的数量，每感染一个减1
// 队列为空时 判断新鲜橘子数量是否为0
func orangesRotting(grid [][]int) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])

	dx := []int{1, -1, 0, 0}
	dy := []int{0, 0, 1, -1}
	count := 0
	for {
		flag := false
		for i := 0; i < rows; i++ {
			for j := 0; j < cols; j++ {
				if grid[i][j] == 3 {
					grid[i][j] = 2
				}
			}
		}
		for i := 0; i < rows; i++ {
			for j := 0; j < cols; j++ {
				if grid[i][j] == 2 {
					for k := 0; k < 4; k++ {
						x := i + dx[k]
						y := j + dy[k]

						if 0 <= x && x < rows && 0 <= y && y < cols && grid[x][y] == 1 {
							grid[x][y] = 3 // 置为3而不是2，为了不影响当前这一分钟的传播
							flag = true
						}
					}
				}
			}
		}

		if !flag {
			for i := range grid {
				for j := range grid[i] {
					if grid[i][j] == 1 {
						return -1
					}
				}
			}
			return count
		}
		count++
	}
}
// @lc code=end

