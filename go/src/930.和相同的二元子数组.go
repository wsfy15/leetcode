/*
 * @lc app=leetcode.cn id=930 lang=golang
 *
 * [930] 和相同的二元子数组
 */

// @lc code=start
func numSubarraysWithSum(nums []int, goal int) int {
	m := map[int]int{}
	m[0] = 1
	res, cur := 0, 0
	for _, num := range nums {
		cur += num
		if v, ok := m[cur-goal]; ok {
			res += v
		}
		m[cur]++
	}

	return res
}

// @lc code=end

