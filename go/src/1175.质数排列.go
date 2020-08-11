/*
 * @lc app=leetcode.cn id=1175 lang=golang
 *
 * [1175] 质数排列
 */

// @lc code=start
func numPrimeArrangements(n int) int {
	primeNum := countPrimes(n)
	res := 1 // 质数的全排列 X 非质数全排列
	for i := 1; i <= primeNum; i++ {
		res = (res * i) % 1_000_000_007
	}
	for i := 1; i <= n-primeNum; i++ {
		res = (res * i) % 1_000_000_007
	}

	return res
}

func countPrimes(n int) int {
	count := 0
	isNotPrime := make([]byte, n)
	for i := 2; i < n; i++ {
		if isNotPrime[i] == 0 {
			count++
			// 这里可以从i*i开始
			// 比如 n = 25，i = 4 时算法会标记 4 × 2 = 8，4 × 3 = 12 等等数字，但是这两个数字已经被 i = 2 和 i = 3 的 2 × 4 和 3 × 4 标记了。
			for j := i * i; j < n; j += i {
				// 标记每一个非质数
				isNotPrime[j] = 1
			}
		}
	}
	return count
}

// @lc code=end

