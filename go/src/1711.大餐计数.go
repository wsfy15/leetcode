/*
 * @lc app=leetcode.cn id=1711 lang=golang
 *
 * [1711] 大餐计数
 */

// @lc code=start
func countPairs(deliciousness []int) int {
	m := map[int]int{}
	var mod int = 1e9 + 7

	res := 0
	for _, v := range deliciousness {
		// 0 <= deliciousness[i] <= 2^20  两数之和最大值为2^21
		for i := 0; i <= 21; i++ {
			if v <= 1<<i {
				if cnt, ok := m[1<<i-v]; ok {
					res = (res + cnt) % mod
				}
			}
		}
		m[v]++
	}

	return res
}

// @lc code=end

