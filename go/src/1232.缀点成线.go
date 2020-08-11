/*
 * @lc app=leetcode.cn id=1232 lang=golang
 *
 * [1232] 缀点成线
 */

// @lc code=start
func checkStraightLine(coordinates [][]int) bool {
	n := len(coordinates)
	if n <= 2 {
		return true
	}

	for i := 2; i < n; i++ {
		if (coordinates[i][0]-coordinates[0][0])*(coordinates[1][1]-coordinates[0][1]) != (coordinates[i][1]-coordinates[0][1])*(coordinates[1][0]-coordinates[0][0]) {
			return false
		}
	}

	return true
}

// @lc code=end

