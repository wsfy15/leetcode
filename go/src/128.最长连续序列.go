/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] 最长连续序列
 */

// @lc code=start
func longestConsecutive(nums []int) int {
	m := map[int]int{} // 记录key所在的连续序列的长度
	maxLen := 0
	for _, num := range nums {
		if _, ok := m[num]; !ok {
			left, _ := m[num - 1]
			right, _ := m[num + 1]
			m[num] = left + right + 1
			m[num - left] = m[num]
			m[num + right] = m[num]
			maxLen = max(maxLen, m[num])
		}
	}
	return maxLen
}

func max(a ...int) int {
	res := a[0]
	for _, v := range a {
		if res < v {
			res = v
		}
	}
	return res
}
// @lc code=end

