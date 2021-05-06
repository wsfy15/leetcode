/*
 * @lc app=leetcode.cn id=403 lang=golang
 *
 * [403] 青蛙过河
 */

// @lc code=start
func canCross(stones []int) bool {
	m := map[int]bool{} // 记录子问题是否遇到过
	return dfs(stones, 0, 0, m)
}

func dfs(stones []int, index, k int, m map[int]bool) bool {
	key := index*1000 + k
	if m[key] {
		return false
	}
	m[key] = true

	for i := index + 1; i < len(stones); i++ {
		distance := stones[i] - stones[index]
		if distance > k+1 {
			break
		}
		if distance < k-1 {
			continue
		}

		if dfs(stones, i, distance, m) {
			return true
		}
	}

	return index == len(stones)-1
}

// @lc code=end

