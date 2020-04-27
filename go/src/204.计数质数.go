/*
 * @lc app=leetcode.cn id=204 lang=golang
 *
 * [204] 计数质数
 */

// @lc code=start
func countPrimes(n int) int { 
  count := 0
  isPrime := make([]byte, n)
  for i := 2; i < n; i++ {
    if isPrime[i] == 0 {
      count++
      // 这里可以从i*i开始
      // 比如 n = 25，i = 4 时算法会标记 4 × 2 = 8，4 × 3 = 12 等等数字，但是这两个数字已经被 i = 2 和 i = 3 的 2 × 4 和 3 × 4 标记了。
      for j := i*i; j < n; j += i {
        // 标记每一个非质数
        isPrime[j] = 1
      }
    }
  }
  return count
}
// @lc code=end

