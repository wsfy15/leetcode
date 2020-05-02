/*
 * @lc app=leetcode.cn id=202 lang=golang
 *
 * [202] 快乐数
 */

// @lc code=start
func isHappy(n int) bool {
	m := map[int]bool{} // 记录求解过程中n的值，如果重复出现，说明进入循环，不可能是快乐数
	m[n] = true
	for n != 1 {
		cur := 0
		for n > 0 {
			tmp := n % 10
			cur = cur + tmp * tmp
			n /= 10
		}
		n = cur
		if _, ok := m[n]; ok {
			return false
		}
		m[n] = true
	}

	return true
}
// @lc code=end

