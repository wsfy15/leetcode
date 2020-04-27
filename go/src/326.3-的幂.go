/*
 * @lc app=leetcode.cn id=326 lang=golang
 *
 * [326] 3的幂
 */

// @lc code=start
func isPowerOfThree(n int) bool {
  // 1162261467 是3在int范围内最大的幂 
  return n > 0 && 1162261467 % n == 0
}
// @lc code=end

