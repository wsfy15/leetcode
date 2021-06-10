/*
 * @lc app=leetcode.cn id=525 lang=golang
 *
 * [525] 连续数组
 */

// @lc code=start
func findMaxLength(nums []int) int {
	cur, res := 0, 0
	m := map[int]int{}
	m[0] = -1
	for i, num := range nums {
		if num == 1 {
			cur++
		} else {
			cur--
		}

		if v, ok := m[cur]; ok {
			res = max(res, i-v)
		} else {
			m[cur] = i
		}
	}

	return res
}

func max(a ...int) int {
	ans := a[0]
	for _, v := range a {
		if v > ans {
			ans = v
		}
	}
	return ans
}

// @lc code=end

