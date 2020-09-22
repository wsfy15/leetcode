/*
 * @lc app=leetcode.cn id=841 lang=golang
 *
 * [841] 钥匙和房间
 */

// @lc code=start
func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)
	visited := make([]bool, n)
	dfs(visited, rooms, 0)
	for _, v := range visited {
		if !v {
			return false
		}
	}
	return true
}

func dfs(visited []bool, rooms [][]int, index int) {
	visited[index] = true
	for _, v := range rooms[index] {
		if !visited[v] {
			dfs(visited, rooms, v)
		}
	}
}

// @lc code=end

