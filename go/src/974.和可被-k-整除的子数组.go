/*
 * @lc app=leetcode.cn id=974 lang=golang
 *
 * [974] 和可被 K 整除的子数组
 */

// @lc code=start
func subarraysDivByK(A []int, K int) int {
	res, n := 0, len(A)
	// kcnt记录 前缀和%K 的个数
	sum, kcnt := make([]int, n + 1), make([]int, K)
	kcnt[0] = 1 // 余数为0的计数初始化为1，当sum[i]为0时，此时就有一个子数组了
	sum[0] = 0
	for i := 1; i <= n; i++ {
		sum[i] = sum[i - 1] + A[i - 1]
		tmp := sum[i] % K
		if tmp < 0 {
			tmp += K
		}
		kcnt[tmp]++
	}
	for i := 0; i < K; i++ {
		// 只要有两个索引的位置kcnt值相同，说明他们之间的子数组是K的倍数
		res += (kcnt[i] * (kcnt[i] - 1)) >> 1
	}

	return res
}
// @lc code=end

