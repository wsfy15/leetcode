/*
 * @lc app=leetcode.cn id=1346 lang=golang
 *
 * [1346] 检查整数及其两倍数是否存在
 */

// @lc code=start
func checkIfExist(arr []int) bool {
	m := map[int]bool{}
	for _, num := range arr {
		if _, ok := m[num]; ok {
			return true
		}
		if num&1 == 0 {
			m[num>>1] = true
		}
		m[num<<1] = true
	}
	return false
}

// @lc code=end

