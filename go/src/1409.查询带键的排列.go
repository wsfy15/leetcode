/*
 * @lc app=leetcode.cn id=1409 lang=golang
 *
 * [1409] 查询带键的排列
 */

// @lc code=start
func processQueries(queries []int, m int) []int {
	nums := make([]int, m)
	for i := 0; i < m; i++ {
		nums[i] = i + 1
	}

	res := make([]int, len(queries))
	for i, query := range queries {
		for j := 0; j < m; j++ {
			if nums[j] == query {
				res[i] = j
				break
			}
		}

		if res[i] > 0 {
			copy(nums[1:res[i]+1], nums[:res[i]])
			nums[0] = query
		}
	}

	return res
}

// @lc code=end

