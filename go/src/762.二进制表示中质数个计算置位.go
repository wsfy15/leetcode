/*
 * @lc app=leetcode.cn id=762 lang=golang
 *
 * [762] 二进制表示中质数个计算置位
 */

// @lc code=start
func countPrimeSetBits(L int, R int) int {
	// 从0到20范围内的质数位置为1
	flag, _ := strconv.ParseInt("10100010100010101100", 2, 32)
	res := 0
	for i := L; i <= R; i++ {
		if 1 << countOneBit(i) & flag > 0 {
			res++
		}
	}
	return res
}

func countOneBit(num int) int {
	count := 0
	for num > 0 {
		if num & 1 > 0 {
			count++
		}
		num >>= 1
	}
	return count
}
// @lc code=end

